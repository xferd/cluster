package config

import (
    "github.com/xferd/golib/mono"
    c "github.com/xferd/cluster/controller"
)

func init() {
    mono.SetViewPath("/Users/leeyan/go/src/github.com/xferd/cluster/template/")
    mono.Handle("^\\/$", &c.HomeController{})
}