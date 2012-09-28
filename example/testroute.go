package main

import (
    "yunio/http"
    "yunio/router"
    "yunio/http/controller"
)

func main() {
    http.Start(
        "localhost:8080",
        router.RegRule{}, 
        http.Patterns{
            "user": &controller.UserController{},
            "files": &controller.UserController{}})
}
