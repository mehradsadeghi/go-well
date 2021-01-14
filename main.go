package main

import (
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

	importPart, startIndex, endIndex := extract(string(file))
	beforeImportContent := file[0:startIndex]
	afterImportContent := file[endIndex+1:]
	importPartLines := strings.Split(importPart, "\n")
	packages := make([]string, 0)
	for _, packageName := range importPartLines {
		packageName = strings.TrimSpace(packageName)
		if packageName != "" {
			packages = append(packages, packageName)
		}
	}

	importablePackages := "    " + strings.Join(packages, "\n    ")
	importPart = "import (\n" + importablePackages + "\n)"

	f, _ := os.OpenFile(FileName, os.O_RDWR, os.ModePerm)
	defer f.Close()

	ip := append(beforeImportContent, []byte(importPart)...)
	ip = append(ip, afterImportContent...)
	ioutil.WriteFile(
		FileName,
		ip,
		0666,
	)
}

func extract(s string) (string, int64, int64) {
	startingPoint := "import ("
	endingPoint := ")"

	startIndex := strings.Index(s, startingPoint)
	if startIndex >= 0 {
		endIndex := strings.Index(s, endingPoint)
		if endIndex >= 0 {
			return s[startIndex+len(startingPoint):endIndex], int64(startIndex), int64(endIndex)
		}
	}
	return "", 0, 0
}
