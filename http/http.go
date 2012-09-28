//bootstrap
package http

import (
    "yunio/router"
    "net/http"
    "yunio/http/controller"
)

func Start(address string) {
    var h handler
    http.ListenAndServe(address, h)
}

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    entry := router.Entry{
                R: req,
                W: w,
                Rule: router.RegRule{},
                Controllers: map[string]router.AbstractController{"user": &controller.UserController{}}}
    router.Route(entry)
}

