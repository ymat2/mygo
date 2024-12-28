package test

import (
    "testing"

	"github.com/ymat2/mygo/cmd"
)

func TestHello(t *testing.T) {
   name := "World"
   want := "Hello, World!"
   if got := cmd.Hello(name); got != want {
       t.Errorf("Hello(%q) = %q; want %q", name, got, want)
   }
}
