package main

import (
    "fmt"
    "os"

    f "flag"

    testing "google.golang.org/protobuf/testing/prototest"
    "github.com/andybalholm/cascadia"
    "github.com/gin-gonic/gin"
)

func helper() {
    cascadia.Parse("A")
    gin.New()
    stupidEnum := testing.Enum{}
    fmt.Println(stupidEnum)
    f.Args()
    os.Exit(1)
}