package main

import (
        "fmt"
        "time"

        "github.com/VimleshS/golang-pipeline/ctx"
        "github.com/VimleshS/golang-pipeline/done"
        "github.com/VimleshS/golang-pipeline/simple"
)

func main() {
        fmt.Println("------------with simple channel  ----------")
        simple.Test()
        fmt.Println("------------with done channel  ----------")
        done.Test()
        fmt.Println("------------with context Done channel  ----------")
        ctx.Test()

        <-time.After(3 * time.Second)
}
