package main

import (
	"io/ioutil"
	"net"
	"os"
	"sort"
	"strings"
)

const FileName = "source.go"

func main() {
	well(FileName)
}

func well(fileName string) {
	file, _ := ioutil.ReadFile(fileName)

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

		// is aliased
		alias := ""
		if strings.Contains(packageName, " ") {
			explodedByWhiteSpace := strings.Split(packageName, " ")
			alias = explodedByWhiteSpace[0]
			packageName = explodedByWhiteSpace[1]
		}

		if strings.Contains(packageName, "/") {
			_, err := net.LookupHost(strings.Split(packageName, "/")[0])
			packageName = "\"" + packageName + "\""
			if len(alias) != 0 {
				packageName = alias + " " + packageName
			}
			if err != nil {
				builtInPackages = append(builtInPackages, packageName)
			} else {
				externalPackages = append(externalPackages, packageName)
			}
		} else {
			packageName = "\"" + packageName + "\""
			if len(alias) != 0 {
				packageName = alias + " " + packageName
			}
			builtInPackages = append(builtInPackages, packageName)
		}
	}

	sort.Strings(builtInPackages)
	sort.Strings(externalPackages)

	temp := ""
	for _, line := range builtInPackages {
		temp = temp + "    " + line + "\n"
	}
	temp = temp + "\n"
	for _, line := range externalPackages {
		temp = temp + "    " + line + "\n"
	}

	importPart = "import (\n" + temp + ")"

	f, _ := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	defer f.Close()

	var ip []byte
	ip = append(ip, beforeImportContent...)
	ip = append(ip, []byte(importPart)...)
	ip = append(ip, afterImportContent...)
	ioutil.WriteFile(fileName, ip, 0666)
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
