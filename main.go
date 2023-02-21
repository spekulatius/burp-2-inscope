package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
)

func main() {
    content, err := os.ReadFile(os.Args[1])

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))

}
