package view

type View interface {
    Assign(k string, v interface{})
    Display(tpl string)
}
