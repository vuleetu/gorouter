package model

type Driver interface {
    New(name Object, source string) (Model, error)
}
