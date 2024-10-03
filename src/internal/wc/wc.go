package wc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func ProcessFile(filename string, lineFlag, wordFlag, charFlag bool, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	if lineFlag {
		countLines(file, filename)
	} else if wordFlag {
		countWords(file, filename)
	} else if charFlag {
		countCharacters(file, filename)
	}
}

func countLines(file *os.File, filename string) {
	var lineCount int
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file %s: %v\n", filename, err)
			break
		}

		if n == 0 {
			break
		}

		lineCount++

		if buffer[n-1] == '\n' {
			lineCount++
		}
	}

	fmt.Printf("%d\t%s\n", lineCount-1, filename)
}

func countWords(file *os.File, filename string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}
	fmt.Printf("%d\t%s\n", wordCount, filename)
}

func countCharacters(file *os.File, filename string) {
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}
	charCount := len([]rune(string(content)))
	fmt.Printf("%d\t%s\n", charCount, filename)
}
