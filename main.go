package main

import (
    "flag"
    "fmt"
    "log"

    "github.com/ymat2/mygo/cmd"
)

func main() {
    var name string
    flag.StringVar(&name, "name", "", "Your name.")
    flag.Parse()

    msg, err := cmd.Hello(name)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msg)
}
