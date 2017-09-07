package controller

import (
    "github.com/xferd/golib/mono"
    "net/http"
    // "log"
)

type HomeController struct {
    mono.Controller
}

func (c *HomeController)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c.Assign("name", "xferd")
    c.Assign("company", "Lenovo")
    c.Display("home.tpl", w)
}

