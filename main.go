package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

var AVAILABLE_OPTIONS []string = []string{"-l", "-c", "-w", "-m"}

func countBytes(data []byte) int {
	return len(data)
}

func countLines(data []byte) int {
	totalLines := bytes.Count(data, []byte{'\n'})

	if len(data) > 0 && !bytes.HasSuffix(data, []byte{'\n'}) {
		totalLines++
	}

	return totalLines
}

func countWords(data []byte) int {
	return len(bytes.Fields(data))
}

func countCharacters(data []byte) int {
	return len(bytes.Split(data, []byte{}))
}

func getOptionAndInputData() ([]string, string, []byte) {
	var pathString, filename string
	data := bytes.Buffer{}
	options := []string{}

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		_, err := io.Copy(&data, os.Stdin)

		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}
	}

	args := os.Args

	for i := 1; i < len(args); i++ {
		argStr := args[i]

		if strings.HasPrefix(argStr, "-") {
			if !slices.Contains(AVAILABLE_OPTIONS, argStr) {
				fmt.Printf("Invalid option : %s, available options are : %v \n", argStr, AVAILABLE_OPTIONS)
				os.Exit(0)
			}
			options = append(options, argStr)
		} else if pathString == "" {
			pathString = argStr
		}
	}

	if pathString != "" {
		file, err := os.Open(pathString)

		if err == nil {
			filename = pathString
			io.Copy(&data, file)
			file.Close()
		} else {
			filename = "total"
			fmt.Println("ERROR : ", err)
		}
	} else {
		filename = "total"
	}

	return options, filename, data.Bytes()
}

func main() {
	options, filename, data := getOptionAndInputData()

	if len(data) == 0 {
		fmt.Println("ERROR : No data")
		return
	}

	if len(options) == 0 {
		options = append(options, "")
	}

	fmt.Print("\t")
	for _, option := range options {
		switch option {
		case "-l":
			totalLine := countLines(data)
			fmt.Print(totalLine, " ")
		case "-c":
			totalBytes := countBytes(data)
			fmt.Print(totalBytes, " ")
		case "-w":
			totalWords := countWords(data)
			fmt.Print(totalWords, " ")
		case "-m":
			totalChars := countCharacters(data)
			fmt.Print(totalChars, " ")
		default:
			totalLine, totalBytes, totalWords := countLines(data), countBytes(data), countWords(data)
			fmt.Print(totalLine, totalWords, totalBytes, " ")
		}
	}
	fmt.Println(filename)

}
