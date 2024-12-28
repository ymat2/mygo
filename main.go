package main

import (
    "flag"

    "github.com/ymat2/mygo/cmd"
)

func main() {
    var name string
    flag.StringVar(&name, "name", "Yuki", "Your name.")
	cmd.Hello(name)
}
