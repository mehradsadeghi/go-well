package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const FileName = "source.txt"

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

func main() {
	file, _ := ioutil.ReadFile(FileName)

	regex := regexp.MustCompile(`import \(([\s\S]*)\)`)
	importPart := regex.FindAllStringSubmatch(string(file), -1)[0][1]

	parts := strings.Split(importPart, "\n")
	packages := make([]string, 0)
	for _, packageName := range parts {
		if packageName != "" {
			packages = append(packages, strings.TrimSpace(packageName))
		}
	}

	importablePakcages := "    " + strings.Join(packages, "\n    ")
	importPart = "import (\n" + importablePakcages + "\n)"

	ioutil.WriteFile(FileName, []byte(importPart), 0666)
}