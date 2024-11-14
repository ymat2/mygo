package mygo

import "testing"

func TestHello(t *testing.T) {
   name := "World"
   want := "Hello, World!"
   if got := Hello(name); got != want {
       t.Errorf("Hello(%q) = %q; want %q", name, got, want)
   }
}
