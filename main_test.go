package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const TestFileName = "test.txt"

func Test_well(t *testing.T) {

	t.Run("it can make the messy #1 well imported", func(t *testing.T) {
		makeTempFileFilledWith(getMessySourceFileContent1())

		well(TestFileName)

		welledContent, err := ioutil.ReadFile(TestFileName)
		if err != nil {
			_ = fmt.Errorf("failed to load well imported file duo to %s", err.Error())
			os.Exit(1)
		}

		t.Log(welledContent)
		t.Log(getWellSourceFileContent1())

		if bytes.Compare(welledContent, getWellSourceFileContent1()) != 0 {
			t.Error("well importing didn't make the import portion well!")
		}

		tearDown()
	})

	t.Run("it can make the messy #2 well imported", func(t *testing.T) {
		makeTempFileFilledWith(getMessySourceFileContent2())

		well(TestFileName)

		welledContent, err := ioutil.ReadFile(TestFileName)
		if err != nil {
			_ = fmt.Errorf("failed to load well imported file duo to %s", err.Error())
			os.Exit(1)
		}

		if bytes.Compare(welledContent, getWellSourceFileContent2()) != 0 {
			t.Error("well importing didn't make the import portion well!")
		}

		tearDown()
	})
}

func tearDown() {
	_ = os.Remove(TestFileName)
}

func makeTempFileFilledWith(content []byte) {
	if err := ioutil.WriteFile(TestFileName, content, 0666); err != nil {
		_ = fmt.Errorf("failed to create temp file duo to %s", err.Error())
		os.Exit(1)
	}
}

func getMessySourceFileContent1() []byte {
	return []byte(`package main

import (
    "fmt"
    "github.com/andybalholm/cascadia"
)

func someFunction() {}`)
}

func getWellSourceFileContent1() []byte {
	return []byte(`package main

import (
    "fmt"

    "github.com/andybalholm/cascadia"
)

func someFunction() {}`)
}

func getMessySourceFileContent2() []byte {
	return []byte(`package main

import (
    "fmt"
    "os/signal"
    "github.com/andybalholm/cascadia"
)

func someFunction() {}`)
}

func getWellSourceFileContent2() []byte {
	return []byte(`package main

import (
    "fmt"
    "os/signal"

    "github.com/andybalholm/cascadia"
)

func someFunction() {}`)
}
