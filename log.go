package main

import (
    "fmt"
)

func log(t string, i string) {
    DBC.Ping()
    _, err := DBC.Exec(fmt.Sprintf("INSERT INTO log SET type = \"%s\", info = \"%s\", time = NOW()", t, i))
    if (err != nil) {
        fmt.Println(err)
    }
}