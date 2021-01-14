package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func Test_well(t *testing.T) {
	fileName := "test.go"
	content := []byte(`package main

import (
    "fmt"
    "github.com/andybalholm/cascadia"
)

func sample() {
    cascadia.Parse("A")
    fmt.Println("A")
}`)
	ioutil.WriteFile(fileName, content, 0666)

	well(fileName)

	welledContent, _ := ioutil.ReadFile(fileName)

	if bytes.Compare(welledContent, content) != 0 {
		t.Error("Didn't work!")
	}

	os.Remove(fileName)
}
