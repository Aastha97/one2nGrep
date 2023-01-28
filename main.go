package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var searchStr, fileName, line, output string
	var result []string
	var isCaseInSensitivity = flag.Bool("i", false, "Ignore case when searching")
	var isWordMatch = flag.Bool("w", false, "Word match when searching")
	var outputFile = flag.String("o", "", "File to write the matches")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Usage: grep [option(s)...] pattern [file_name]")
		return
	}
	nonFlagValues := flag.Args()
	if strings.Contains(os.Args[len(os.Args)-1], ".txt") {
		searchStr = nonFlagValues[0]
		fileName = nonFlagValues[1]
	} else {
		searchStr = os.Args[len(os.Args)-1]
	}

	if fileName == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				line = scanner.Text()
			}
			output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)
			if output != "" {
				fmt.Println(output + "\n")
			}
		}
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line = scanner.Text()
			output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)
			if output != "" {
				result = append(result, output)
			}

		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if outputFile != nil {
		writeFile(result, *outputFile)
	} else {
		finalResult(strings.Join(result, "\n"))
	}

}
