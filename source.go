package main

import (
    "fmt"
    "os"
    "os/signal"
    f "flag"

    "github.com/andybalholm/cascadia"
    a "github.com/andybalholm/cascadia/fuzz"
    "github.com/gin-gonic/gin"
)

func justForIgnoringErrors() {
    cascadia.Parse("A")
    gin.New()
    signal.Ignore()
    //stupidEnum := testing.Enum{}
    //fmt.Println(stupidEnum)
    a.Fuzz([]byte{})
    fmt.Println("A")
    f.Args()
    os.Exit(1)
}