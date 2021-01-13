package main

import (
    "fmt"
    "os"

    "github.com/andybalholm/cascadia"
    "github.com/gin-gonic/gin"
)

func helper() {
    cascadia.Parse("A")
    fmt.Println("A")
    gin.New()
    os.Exit(1)
}