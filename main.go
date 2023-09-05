package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	fmt.Println("Input file path: ")
	var fileName string
	if _, err := fmt.Scanln(&fileName); err != nil {
		log.Fatal(err)
	}
	if err := ReadFileAndProcess(fileName); err != nil {
		log.Fatal(err)
	}
}

func ReadFileAndProcess(filename string) error {
	byteArray, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lineArray := bytes.Split(byteArray, []byte("\n"))

	// fmt.Println(string(byteArray))
	for i, v := range lineArray {
		if bytes.Contains(v, []byte("type")) && bytes.Contains(v, []byte("struct")) {
			structName := bytes.Split(v, []byte(" "))[1]
			j := i + 1
			for j < len(lineArray) {
				if bytes.Contains(lineArray[j], []byte("}")) {
					break
				}
				newlineWithoutExtraSpace := getNewLine(string(lineArray[j]))

				newLine := strings.Split(newlineWithoutExtraSpace, " ")

				if len(newLine) >= 2 && newLine[0] != "//" && newLine[1] != "struct" && newLine[1] != "func" {
					tagName := newLine[0]
					tagType := newLine[1]
					get, set := createFunction(tagName, tagType, structName)
					WriteInFile(filename, get)
					WriteInFile(filename, set)
				}
				j++
			}
		}
	}
	return nil
}

func getNewLine(line string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(line, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")
	return final
}

func WriteInFile(fileName string, res []byte) error {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 064)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(string(res))
	if err != nil {
		return err
	}

	return nil
}

func createFunction(tagName, tagType string, structName []byte) ([]byte, []byte) {

	smallTagName := strings.ToLower(tagName)

	funcTitle := cases.Title(language.English, cases.NoLower).String(tagName)
	smallCaseOfStruct := strings.ToLower(string(structName[0]))
	getfuncDefinition := "\nfunc (" + smallCaseOfStruct + " *" + string(structName) + ") Get" + funcTitle + "() " + string(tagType) + " {" + "\n\t" + "return " + smallCaseOfStruct + "." + strings.TrimSpace(string(tagName)) + "\n" + "} \n"
	setfuncDefinition := "\nfunc (" + smallCaseOfStruct + " *" + string(structName) + ") Set" + funcTitle + "(" + strings.TrimSpace(smallTagName) + " " + string(tagType) + ") {" + "\n\t" + smallCaseOfStruct + "." + strings.TrimSpace(string(tagName)) + " = " + strings.TrimSpace(smallTagName) + "\n" + "} \n"

	return []byte(getfuncDefinition), []byte(setfuncDefinition)

}
