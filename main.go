package main

import (
	"fmt"
	"os"
	"strings"
)

func countBytes(data []byte) int {
	return len(data)
}

func countLines(data []byte) int {
	stringData := strings.Split(string(data), "\n")
	return len(stringData)
}

func countWords(data []byte) int {
	return len(strings.Fields(string(data)))
}

func countCharacters(data []byte) int {
	return len(strings.Split(string(data), ""))
}

func main() {
	args := os.Args

	var option, filePath string

	if strings.HasPrefix(args[1], "-") {
		option = args[1]
		filePath = strings.Join(args[2:], " ")
	} else {
		filePath = strings.Join(args[1:], " ")
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		data = []byte(filePath)
	}

	switch option {
	case "-l":
		totalLine := countLines(data)
		fmt.Println(totalLine, filePath)
	case "-c":
		totalBytes := countBytes(data)
		fmt.Println(totalBytes, filePath)
	case "-w":
		totalWords := countWords(data)
		fmt.Println(totalWords, filePath)
	case "-m":
		totalChars := countCharacters(data)
		fmt.Println(totalChars, filePath)
	default:
		totalLine, totalBytes, totalWords := countLines(data), countBytes(data), countWords(data)
		fmt.Println(totalLine, totalWords, totalBytes, filePath)
	}

}
