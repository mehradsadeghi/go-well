package main

import (
	"io/ioutil"
	"net"
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
	builtInPackages := make([]string, 0)
	externalPackages := make([]string, 0)
	for _, packageName := range importPartLines {
		packageName = strings.TrimSpace(packageName)
		packageName = strings.ReplaceAll(packageName, "\"", "")

		if packageName == "" {
			continue
		}

		if strings.Contains(packageName, "/") {
			_, err := net.LookupHost(strings.Split(packageName, "/")[0])
			if err != nil {
				builtInPackages = append(builtInPackages, "\"" + packageName + "\"")
			} else {
				externalPackages = append(externalPackages, "\"" + packageName + "\"")
			}
		} else {
			builtInPackages = append(builtInPackages, "\"" + packageName + "\"")
		}
	}

	temp := ""
	for _, line := range builtInPackages {
		temp = temp + "    " + line + "\n"
	}
	temp = temp + "\n"
	for _, line := range externalPackages {
		temp = temp + "    " + line + "\n"
	}

	importPart = "import (\n" + temp + ")"

	f, _ := os.OpenFile(FileName, os.O_RDWR, os.ModePerm)
	defer f.Close()

	var ip []byte
	ip = append(ip, beforeImportContent...)
	ip = append(ip, []byte(importPart)...)
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
