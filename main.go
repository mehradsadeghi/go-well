package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const FileName = "source.go"

/*func main() {
	file, _ := ioutil.ReadFile("source.txt")

	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return
	}

	ioutil.WriteFile(
		"source.txt",
		[]byte(regex.ReplaceAllString(string(file), "\n")),
		0666,
	)
}*/

//regex := regexp.MustCompile(`import \(([\s\S]*)\)`)
//importPart := regex.FindAllStringSubmatch(string(file), -1)[0][1]

func main() {
	file, _ := ioutil.ReadFile(FileName)

	importPart, startIndex, _ := match(string(file))
	parts := strings.Split(importPart, "\n")
	packages := make([]string, 0)
	for _, packageName := range parts {
		if packageName != "" {
			packages = append(packages, strings.TrimSpace(packageName))
		}
	}

	importablePackages := "    " + strings.Join(packages, "\n    ")
	importPart = "import (\n" + importablePackages + "\n)"

	fmt.Println(importPart)

	f, _ := os.OpenFile(FileName, os.O_RDWR, os.ModePerm)
	defer f.Close()
	f.WriteAt([]byte{'M', 'E', 'H', 'R', 'A', 'D'}, startIndex)
	// ioutil.WriteFile(FileName, []byte(importPart), 0666)

	// replace the import part with a placeholder
	// and then replace the final string with the placeholder
}

func match(s string) (string, int64, int) {
	startIndex := strings.Index(s, "import (")
	if startIndex >= 0 {
		endIndex := strings.Index(s, ")")
		if endIndex >= 0 {
			return s[startIndex+8:endIndex], int64(startIndex), endIndex
		}
	}
	return "", 0, 0
}
