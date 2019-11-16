package lookup

import (
    "fmt"
    "strings"
)

import (
    "github.com/eyedeekay/sam3"
)

func Sane(addr string) error {
    if ! strings.HasSuffix(addr, ".i2p") {
        return fmt.Errorf("Error! This is not an i2p address")
    }
    return nil
}

func Lookup(addr string) (string, error){
    if err := Sane(addr); err != nil {
        return "", err
    }
    sam, err := sam3.NewSAM("127.0.0.1:7656")
    if err != nil {
        return "", err
    }
    defer sam.Close()
    raddrkeys, err := sam.Lookup(addr)
    if err != nil {
        return "", err
    }
    raddr := raddrkeys.Base64()
    return raddr, nil
}
