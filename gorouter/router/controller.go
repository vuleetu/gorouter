// Package router provides ...
package router

import (
    "net/http"
    "fmt"
)

type AbstractController interface {
    Error(int)
    SetW(w http.ResponseWriter)
    SetR(r *http.Request)
    GetW() http.ResponseWriter
    GetR() *http.Request
}

type Controller struct {
    w http.ResponseWriter
    r *http.Request
}

func (c *Controller) Error(code int){
    fmt.Fprintf(c.w, "The http error code is %d", code)
}

func (c *Controller) SetW(w http.ResponseWriter){
    c.w = w
}

func (c *Controller) SetR(r *http.Request){
    c.r = r
}


func (c *Controller) GetW() http.ResponseWriter{
    return c.w
}

func (c *Controller) GetR() *http.Request{
    return c.r
}
