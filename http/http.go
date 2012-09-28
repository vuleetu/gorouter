//bootstrap
package http

import (
    "yunio/router"
    "net/http"
)

type Patterns map[string]router.AbstractController

func Start(address string, rule router.Rule, patterns Patterns) {
    h := handler{rule, patterns}
    http.ListenAndServe(address, h)
}

type handler struct{
    rule router.Rule
    patterns Patterns
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    entry := router.Entry{
                R: req,
                W: w,
                Rule: h.rule,
                Controllers: h.patterns}
    router.Route(entry)
}

