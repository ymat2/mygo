package main

import (
    "flag"
    "fmt"

    "github.com/ymat2/mygo/cmd"
)

func main() {
    var name string
    flag.StringVar(&name, "name", "Yuki", "Your name.")
    flag.Parse()

    fmt.Println(cmd.Hello(name))
}
