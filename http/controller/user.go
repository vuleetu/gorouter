package controller

import (
    "yunio/router"
    "yunio/view"
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
    tpl := view.NewDjangoTpl(c.GetW())
    tpl.Assign("title", "Yunio")
    tpl.Assign("name", "view based on django")
    tpl.Display("example.html")
}

