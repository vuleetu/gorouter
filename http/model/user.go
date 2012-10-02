package model

import (
    m "yunio/gorouter/model"
    "log"
)

type User struct {
    id uint64
    username, email string
}

func New(source string ) m.Model {
    model, err := m.New("mysql", User{},source)
    if err != nil {
        log.Println(err.Error())
        return nil
    }
    return model
}
