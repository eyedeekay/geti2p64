package main

import (
    "os"
    "fmt"
)

import (
    "github.com/eyedeekay/geti2p64"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please provide a single i2p hostname or b32 as an argument")
    }
    base64, err := lookup.Lookup(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
    if base64 == "" {
        fmt.Println("Hostname or b32 not found")
    }
    fmt.Println(base64)
}
