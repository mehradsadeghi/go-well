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
		assertTestFileEqualsTo(t, getWellSourceFileContent1())
		tearDown()
	})

	t.Run("it can make the messy #2 well imported", func(t *testing.T) {
		makeTempFileFilledWith(getMessySourceFileContent2())
		well(TestFileName)
		assertTestFileEqualsTo(t, getWellSourceFileContent2())
		tearDown()
	})

	t.Run("it can make the messy #3 well imported", func(t *testing.T) {
		makeTempFileFilledWith(getMessySourceFileContent3())
		well(TestFileName)
		assertTestFileEqualsTo(t, getWellSourceFileContent3())
		tearDown()
	})

	t.Run("it can make the messy and unsorted #1 well sorted and imported", func(t *testing.T) {
		makeTempFileFilledWith(getUnSortedMessySourceFileContent1())
		well(TestFileName)
		assertTestFileEqualsTo(t, getWellSortedSourceFileContent1())
		tearDown()
	})

	t.Run("it can make the messy and unsorted #2 well sorted and imported", func(t *testing.T) {
		makeTempFileFilledWith(getUnSortedMessySourceFileContent2())
		well(TestFileName)
		assertTestFileEqualsTo(t, getWellSortedSourceFileContent2())
		tearDown()
	})
}

func assertTestFileEqualsTo(t *testing.T, content []byte) {
	t.Helper()
	welledContent, err := ioutil.ReadFile(TestFileName)
	if err != nil {
		_ = fmt.Errorf("failed to load well imported file duo to %s", err.Error())
		os.Exit(1)
	}

	if bytes.Compare(welledContent, content) != 0 {
		t.Error("well importing didn't make the import portion well!")
	}
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

func getMessySourceFileContent3() []byte {
	return []byte(`package main

import (
    "fmt"
    fuzz "github.com/andybalholm/cascadia/fuzz"
    "github.com/andybalholm/cascadia"
    "os/signal"
    f "flag"
    "os"
)

func someFunction() {}`)
}

func getWellSourceFileContent3() []byte {
	return []byte(`package main

import (
    f "flag"
    "fmt"
    "os"
    "os/signal"

    "github.com/andybalholm/cascadia"
    fuzz "github.com/andybalholm/cascadia/fuzz"
)

func someFunction() {}`)
}

func getUnSortedMessySourceFileContent1() []byte {
	return []byte(`package main

import (
    "os"
    "fmt"
)

func someFunction() {}`)
}

func getWellSortedSourceFileContent1() []byte {
	return []byte(`package main

import (
    "fmt"
    "os"

)

func someFunction() {}`)
}

func getUnSortedMessySourceFileContent2() []byte {
	return []byte(`package main

import (
    "fmt"
    "os"
    "github.com/example"
    "github.com/aexample"
)

func someFunction() {}`)
}

func getWellSortedSourceFileContent2() []byte {
	return []byte(`package main

import (
    "fmt"
    "os"

    "github.com/aexample"
    "github.com/example"
)

func someFunction() {}`)
}

// todo test cases: imports without packages consist of `/`
// todo test cases: imports without builtin packages