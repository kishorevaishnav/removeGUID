package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	// Replace "oldString" with "newString" in the file.
	// Replace this with your desired strings.
	oldString := `\S{8}-(\S{4}-){3}\S{12}-\S{8}(:\S{8}){0,1}`
	newString := ``

	filePath := os.Args[1] // Replace this with your file path.

	// Read the content of the file.
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Convert UTF-16 LE bytes to a string.
	contentString, err := decodeUTF16LE(content)
	if err != nil {
		fmt.Println("Error decoding UTF-16 LE:", err)
		os.Exit(1)
	}

	// Create a regular expression to find the oldString.
	regex := regexp.MustCompile(oldString)

	// Replace the oldString with the newString using the regex.
	updatedContent := regex.ReplaceAllString(contentString, newString)

	// Convert the updated content string back to UTF-16 LE bytes.
	updatedContentBytes, err := encodeUTF16LE(updatedContent)
	if err != nil {
		fmt.Println("Error encoding UTF-16 LE:", err)
		os.Exit(1)
	}

	// Replace CRLF with LF (if desired).
	updatedContentBytes = []byte(strings.Replace(string(updatedContentBytes), "\r\n", "\n", -1))

	suffix := "_fixed"
	baseName := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	extension := filepath.Ext(filePath)

	// Append the suffix to the base name
	newFileName := baseName + suffix + extension

	// Write the updated content back to the file.
	err = ioutil.WriteFile(newFileName, updatedContentBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("String replaced successfully.")
}

func decodeUTF16LE(b []byte) (string, error) {
	utf16Decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	utf8Reader := transform.NewReader(strings.NewReader(string(b)), utf16Decoder.NewDecoder())
	decoded, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func encodeUTF16LE(s string) ([]byte, error) {
	utf16Encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	utf16Bytes, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(s), utf16Encoder.NewEncoder()))
	if err != nil {
		return nil, err
	}
	return utf16Bytes, nil
}
