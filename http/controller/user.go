package controller

import (
    /*"yunio/gorouter"
    "yunio/view"*/
    "fmt"
)

type UserController struct {
    router.Controller
}

func (c *UserController) TestAction() {
    fmt.Fprint(c.GetW(), "test action triggered")
}

func (c *UserController) HelloAction() {
    fmt.Fprint(c.GetW(), "hello action triggered")
    fmt.Fprintf(c.GetW(), "Host is %s", c.GetR().Host)
}

func (c *UserController) TemplateAction() {
    defer db.Close()
    rows, err1 := db.Query("SELECT path, name FROM user_files LIMIT 10")
    if err1 != nil {
        fmt.Fprintf(c.GetW(), "Query failed, error is %s", err1.Error())
        return
    }

    for rows.Next() {
        var path, name string
        err = rows.Scan(&path, &name)
        if err != nil {
            fmt.Fprintf(c.GetW(), "Get rows failed, error is %s", err.Error())
            return
        }
        fmt.Println(path, name)
    }

    tpl := view.NewDjangoTpl(c.GetW())
    tpl.Assign("title", "Yunio")
    tpl.Assign("name", "view based on django")
    tpl.Display("example.html")
}

