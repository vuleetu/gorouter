package db

import (
    "reflect"
    "strings"
    "fmt"
    "errors"
    "github.com/ziutek/mymysql/mysql"
    "io"
    "time"
    "log"
)

type Table struct {
    fields []Field //fileds of table
    name string //table name
}

type Field struct {
    Name string
    Key  string
    Type reflect.Type
}

type Object struct{
    Type reflect.Type
    Value reflect.Value
}

type Model struct {
    tbl *Table
    obj *Object
    conn mysql.Conn
}

type Data interface{}
type Datas []Data

func (m *Model) GetAll() (Datas, error) {
    fields := make([]string, 0, 10)
    for i := 0; i < len(m.tbl.fields); i++ {
        fields = append(fields, m.tbl.fields[i].Name)
    }
    sql := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ","), m.tbl.name)
    //res, err := m.conn.Start(sql)
    stmt, err := m.conn.Prepare(sql)
    if err != nil {
        return nil, err
    }

    res, err := stmt.Run()
    if err != nil {
        return nil, err
    }

    objs := make([]Data, 0, 100) //100 first
    row := res.MakeRow()
    for {
        err = res.ScanRow(row)
        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }

        objs = append(objs, MakeData(m.obj.Type, row))
    }

    return objs, nil
}

func (m *Model) Save() error {
    datas := make([]interface{}, 0, 10)
    columns := make([]string, 0, 10)
    length := 0
    for i := 0; i < len(m.tbl.fields); i++ {
        if m.tbl.fields[i].Key != "true" {
            length++
            datas = append(datas, m.obj.Value.Field(i).Interface())
            columns = append(columns, m.tbl.fields[i].Name)
        }
    }
    log.Println(datas, columns, length)
    columnStr := strings.Join(columns, ",")
    dataStr := strings.Repeat("?,", length)
    dataStr = dataStr[0:len(dataStr)-1]
    sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", m.tbl.name, columnStr, dataStr)
    log.Println("Sql is ", sql)
    stmt, err := m.conn.Prepare(sql)
    if err != nil {
        return err
    }
    stmt.Bind(datas...)
    res, err := stmt.Run()
    if err != nil {
        return err
    }

    fmt.Println("Affected rows is ", res.AffectedRows())
    return nil
}

func (m *Model) Update() error {
    datas := make([]interface{}, 0, 10)
    columns := make([]string, 0, 10)
    keydatas := make([]interface{}, 0, 10)
    keycolumns := make([]string, 0, 10)
    for i := 0; i < len(m.tbl.fields); i++ {
        if m.tbl.fields[i].Key != "true" {
            datas = append(datas, m.obj.Value.Field(i).Interface())
            columns = append(columns, m.tbl.fields[i].Name)
        } else {
            log.Println("is key")
            keydatas = append(keydatas, m.obj.Value.Field(i).Interface())
            keycolumns = append(keycolumns, m.tbl.fields[i].Name)
        }
    }
    log.Println(datas, columns)
    columnStr := strings.Join(columns, "=?, ")
    log.Println(columnStr)
    columnStr = columnStr[0:len(columnStr)-2]
    keyColumnStr := strings.Join(keycolumns, "=? AND ")
    log.Println(keyColumnStr)
    keyColumnStr = keyColumnStr[0:len(keyColumnStr)-5]
    sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", m.tbl.name, columnStr, keyColumnStr)
    log.Println("Sql is ", sql)
    stmt, err := m.conn.Prepare(sql)
    if err != nil {
        return err
    }
    alldatas := append(datas, keydatas...)
    stmt.Bind(alldatas...)
    res, err := stmt.Run()
    if err != nil {
        return err
    }

    fmt.Println("Affected rows is ", res.AffectedRows())
    return nil
}

func MakeData(Type reflect.Type, row mysql.Row) Data {
    //log.Println(row)
    v := reflect.New(Type).Elem() //Use new to get the pointer
    log.Println("The type of v is ", v.Type())
    length := v.NumField()
    for i := 0; i < length; i++ {
        switch row[i].(type) {
        case []byte: //varchar
            v.Field(i).SetString(string(row[i].([]byte)))
        case string: //string
            v.Field(i).SetString(row[i].(string))
        case uint:
            v.Field(i).SetUint(uint64(row[i].(uint)))
        case uint8:
            v.Field(i).SetUint(uint64(row[i].(uint8)))
        case uint16:
            v.Field(i).SetUint(uint64(row[i].(uint16)))
        case uint32:
            v.Field(i).SetUint(uint64(row[i].(uint32)))
        case uint64:
            v.Field(i).SetUint(uint64(row[i].(uint64)))
        case int:
            v.Field(i).SetInt(int64(row[i].(int)))
        case int8:
            v.Field(i).SetInt(int64(row[i].(int8)))
        case int16:
            v.Field(i).SetInt(int64(row[i].(int16)))
        case int32:
            v.Field(i).SetInt(int64(row[i].(int32)))
        case int64:
            v.Field(i).SetInt(int64(row[i].(int64)))
        case bool:
            v.Field(i).SetBool(row[i].(bool))
        case time.Time:
            v.Field(i).Set(reflect.ValueOf(row[i].(time.Time)))
        }
    }

    return v.Interface()
}

func NewModel(obj interface{}, conn mysql.Conn) *Model {
    tbl, o, err := TableInfo(obj)

    if err != nil {
        log.Println(err.Error())
        return nil
    }

    return &Model{tbl, o, conn}
}

func TableInfo(object interface{}) (*Table, *Object, error) {
    t := reflect.TypeOf(object)

    if t.Kind() != reflect.Struct {
        return nil, nil, errors.New("invalid object")
    }

    fields := make([]Field,0, 100) // is 100 enough?
    length := t.NumField()

    for i := 0; i< length; i++ {
        log.Println("Found field ", t.Field(i).Name, ", Tag is ", t.Field(i).Tag)
        fields = append(fields,
                        Field{
                            strings.ToLower(t.Field(i).Name),
                            t.Field(i).Tag.Get("key"),
                            t.Field(i).Type})
    }

    return &Table{fields, strings.ToLower(t.Name())},
           &Object{t, reflect.ValueOf(object)},
           nil
}

func DataToModel(row []string, table *Table) {
}
