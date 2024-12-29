package cmd

import (
    "errors"
    "fmt"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("No name given. Are you heiße?")
    }

    message := fmt.Sprintf("Hello, %v!", name)
    return message, nil  // nil is type means no error
}
