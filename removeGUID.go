package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		fmt.Println("OR Usage: removeGUID <input_file> <output_file>")
		return
	}
	inputFilePath := os.Args[1]
	// outputFilePath := os.Args[2]

	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("err: ", err)
	}

	newContent, _, _ := transform.String(charmap.Windows1252.NewDecoder(), string(content))

	re := regexp.MustCompile(`\S{8}-(\S{4}-){3}\S{12}-\S{8}(:\S{8}){0,1}`)
	newBContent := re.ReplaceAll([]byte(newContent), []byte(""))

	err = ioutil.WriteFile(appendRGToFile(inputFilePath), newBContent, 0644)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func appendRGToFile(fileNamePath string) string {
	dir, file := filepath.Split(fileNamePath)
	ext := filepath.Ext(file)
	newFile := file[:len(file)-len(ext)] + "_rg" + ext
	newPath := filepath.Join(dir, newFile)
	return newPath
}
