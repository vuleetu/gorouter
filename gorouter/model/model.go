package model

import (
    "errors"
)

var models = make(map[string]Driver)

type Data interface{}
type Datas []Data

type Model interface {
    GetAll() (Datas, error)
    Save() error
    Update() error
    Delete() error
}

func New(drvname string, obj Object, source string) (Model, error) {
    driver, ok := models[drvname]
    if !ok {
        return nil, errors.New("Driver not found")
    }
    return driver.New(obj, source)
}

func Register(t string, driver Driver) {
    models[t] = driver
}
