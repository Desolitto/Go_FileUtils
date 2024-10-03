package parser

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlagsFind() (*bool, *bool, *bool, *string) {
	fileFlag := flag.Bool("f", false, "Find Files")
	dirFlag := flag.Bool("d", false, "Find Directories")
	symlinkFlag := flag.Bool("sl", false, "Find Symbolic links")
	extFlag := flag.String("ext", "", "File extension (only with -f)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a directory.")
		os.Exit(1)
	}

	return fileFlag, dirFlag, symlinkFlag, extFlag
}

func ParseFlags() (lineFlag, wordFlag, charFlag bool) {
	flag.BoolVar(&lineFlag, "l", false, "count lines")
	flag.BoolVar(&wordFlag, "w", false, "count words")
	flag.BoolVar(&charFlag, "m", false, "count characters")
	flag.Parse()

	// Проверка на совместимость флагов
	if (lineFlag && wordFlag) || (lineFlag && charFlag) || (wordFlag && charFlag) {
		fmt.Println("Ошибка: можно использовать только один флаг (-l, -m или -w)")
		os.Exit(1)
	}

	if !lineFlag && !wordFlag && !charFlag {
		wordFlag = true // По умолчанию использовать подсчет слов
	}

	return
}
