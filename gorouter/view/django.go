package view

import (
    "github.com/flosch/pongo"
    "net/http"
)

func NewDjangoTpl(w http.ResponseWriter) *DjangoTpl{
    return &DjangoTpl{make(pongo.Context), w}
}

type DjangoTpl struct {
    ctx pongo.Context
    w http.ResponseWriter
}

func (tpl *DjangoTpl) Assign(k string, v interface{}) {
    tpl.ctx[k] = v
}


func (tpl *DjangoTpl) Display(template string) {
    var tplExample = pongo.Must(pongo.FromFile(template, nil))
    err := tplExample.ExecuteRW(tpl.w, &tpl.ctx)
    if err != nil {
        http.Error(tpl.w, err.Error(), http.StatusInternalServerError)
    }
}
