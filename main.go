package main

import (
    "github.com/xferd/golib/mono"
    _ "github.com/xferd/cluster/config"
)

func main() {
    mono.ListenAndServe(":8081")
}