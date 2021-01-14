package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

const FileName = "source.go"

func main() {
	well(FileName)
}

func well(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	importContents, beforeImportContents, afterImportContents := extractImportContents(string(file))
	if len(importContents) == 0 {
		return fmt.Errorf("there is no import in %s", fileName)
	}

	importLines := normalizeImportLines(importContents)

	builtInPackages := make([]string, 0)
	externalPackages := make([]string, 0)
	for _, packageName := range importLines {
		aliasName := ""
		if isAliased(packageName) {
			aliasName, packageName = extractAliasedPackage(packageName)
		}

		if !strings.Contains(packageName, "/") {
			builtInPackages = appendTo(builtInPackages, packageName, aliasName)
		} else {
			_, err := net.LookupHost(strings.Split(packageName, "/")[0])
			packageName = "\"" + packageName + "\""
			if len(aliasName) != 0 {
				packageName = aliasName + " " + packageName
			}
			if err != nil {
				builtInPackages = append(builtInPackages, packageName)
			} else {
				externalPackages = append(externalPackages, packageName)
			}
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

	importContents = "import (\n" + temp + ")"

	// writing back into the file
	f, _ := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	defer f.Close()

	var ip []byte
	ip = append(ip, beforeImportContents...)
	ip = append(ip, []byte(importContents)...)
	ip = append(ip, afterImportContents...)
	ioutil.WriteFile(fileName, ip, 0666)

	return nil
}

func appendTo(builtInPackages []string, packageName, aliasName string) []string {
	packageName = "\"" + packageName + "\""
	if len(aliasName) != 0 {
		packageName = aliasName + " " + packageName
	}
	builtInPackages = append(builtInPackages, packageName)
	return builtInPackages
}

func extractAliasedPackage(name string) (alias, packageName string) {
	explodedByWhiteSpace := strings.Split(name, " ")
	return explodedByWhiteSpace[0], explodedByWhiteSpace[1]
}

func isAliased(packageName string) bool {
	return strings.Contains(packageName, " ")
}

func normalizeImportLines(importContent string) []string {
	importLines := strings.Split(importContent, "\n")

	normalizedImportLines := make([]string, 0)

	for _, packageName := range importLines {
		packageName = strings.TrimSpace(packageName)
		packageName = strings.ReplaceAll(packageName, "\"", "")

		if packageName == "" {
			continue
		}

		normalizedImportLines = append(normalizedImportLines, packageName)
	}

	return normalizedImportLines
}

func extractImportContents(content string) (importContent, beforeImportContent, afterImportContent string) {
	startsWith := "import ("
	endsWith := ")"

	startOfImport := strings.Index(content, startsWith)
	endOfImport := strings.Index(content, endsWith)
	if startOfImport < 0 || endOfImport < 0 {
		return
	}

	beforeImportContent = content[0:int64(startOfImport)]
	afterImportContent = content[endOfImport+1:]
	importContent = content[startOfImport+len(startsWith): endOfImport]

	return
}
