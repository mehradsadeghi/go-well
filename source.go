package main

import (
    "fmt"
    "os"
    
    "github.com/andybalholm/cascadia"
)
// "github.com/gin-gonic/gin"
// "os/signal"
// f "flag"
// testing "google.golang.org/protobuf/testing/prototest"

func justForIgnoringErrors() {
    cascadia.Parse("A")
    //gin.New()
    //signal.Ignore()
    //stupidEnum := testing.Enum{}
    //fmt.Println(stupidEnum)
    fmt.Println("A")
    //f.Args()
    os.Exit(1)
}