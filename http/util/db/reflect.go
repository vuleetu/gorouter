package db

import (
    "reflect"
    "bytes"
)

func ModelToColumn(model interface{}) string {
    t := reflect.TypeOf(model)
    length := t.NumField()
    buf := bytes.NewBuffer(nil)
    for i := 0; i< length; i++ {
        sf := t.Field(i)
        buf.WriteString(sf.Name)
        if i < length-1 {
            buf.WriteString(",")
        }
    }
    return buf.String()
}

func DataToModel([][]byte, model interface{}) {
    
}
