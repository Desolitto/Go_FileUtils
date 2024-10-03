package main

import (
	"fileutils/internal/parser"
	"fileutils/internal/wc"
	"flag"
	"sync"
)

func main() {
	lineFlag, wordFlag, charFlag := parser.ParseFlags()
	var wg sync.WaitGroup

	// Обработка файлов
	for _, filename := range flag.Args() {
		wg.Add(1)
		go wc.ProcessFile(filename, lineFlag, wordFlag, charFlag, &wg)
	}

	wg.Wait() // Ждем завершения всех горутин
}
