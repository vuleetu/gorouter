package db

import (
    "reflect"
    "strings"
    "fmt"
    "errors"
    "github.com/ziutek/mymysql/mysql"
    "io"
    _ "log"
)

type Table struct {
    fields []string //fileds of table
    name string //table name
}

type Object struct{
    Type reflect.Type
}

type Model struct {
    tbl *Table
    obj *Object
    conn mysql.Conn
}

type Data interface{}
type Datas []Data

func (m *Model) GetAll() (Datas, error) {
    sql := fmt.Sprintf("SELECT %s FROM user_files", strings.Join(m.tbl.fields, ","))
    res, err := m.conn.Start(sql)
    if err != nil {
        return nil, err
    }

    objs := make([]Data, 100) //100 first
    row := res.MakeRow()
    for {
        err = res.ScanRow(row)
        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }
    }
}

func MakeData(Type reflect.Type, row mysql.Row) {
    v := reflect.New(Type) //Use new to get the pointer
    length := v.Len()
    for i := 0; i < length; i++ {
        assign(row[i], v.Type())
    }
    d := v.interface{}
    reflect.TypeOf()
}

func assign(data interface{}, Type reflect.Type) {
    switch Type {
        case reflect.String: return data.(string)
        case reflect.Uint, reflect.Uint8, reflect.Uint16,
    }
}

/*func NewModel(obj interface{}, db *sql.DB) *Model {
    table, err := TableInfo(obj)

    if err != nil {
        log.Println(err.Error())
        return nil
    }

    return &Model{table, obj, db}
}*/

func TableInfo(object interface{}) (*Table, error) {
    t := reflect.TypeOf(object)

    if t.Kind() != reflect.Struct {
        return nil, errors.New("invalid object")
    }

    fields := make([]string, 100) // is 100 enough?
    length := t.NumField()

    for i := 0; i< length; i++ {
        fields[i] = strings.ToLower(t.Field(i).Name)
    }

    return &Table{fields, strings.ToLower(t.Name())}, nil
}

func DataToModel(row []string, table *Table) {
}
