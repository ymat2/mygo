package cmd

import ("testing")

func TestHello(t *testing.T) {
   name := "World"
   want := "Hello, World!"
   if std, err := Hello(name); std != want || err != nil {
       t.Errorf("Hello(%q) = %q; want %q", name, std, want)
   }
}
