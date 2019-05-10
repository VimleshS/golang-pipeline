package main

import (
        "fmt"
        "time"

        "github.com/VimleshS/testPipeline/ctx"
        "github.com/VimleshS/testPipeline/done"
        "github.com/VimleshS/testPipeline/simple"
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
