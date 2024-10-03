package main

import (
	"fileutils/internal/rotate"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	archiveDir := flag.String("a", "", "Directory to store archive files")
	flag.Parse()

	logFiles := flag.Args()
	if len(logFiles) < 1 {
		fmt.Println("Usage: myRotate [-a <archive_directory>] <log_file1> [<log_file2> ...]")
		return
	}

	if *archiveDir == "" {
		*archiveDir = filepath.Dir(logFiles[0])
	}

	if err := os.MkdirAll(*archiveDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", *archiveDir, err)
		return
	}

	var wg sync.WaitGroup

	for _, logFile := range logFiles {
		if filepath.Ext(logFile) != ".log" {
			fmt.Printf("Skipping file %s: only .log files are allowed\n", logFile)
			continue
		}

		wg.Add(1)
		go rotate.ArchiveFile(logFile, *archiveDir, &wg)
	}

	wg.Wait()
}

/*
go build -o myRotate ./cmd/myRotate
./myRotate -a testdata/archive testdata/test04/logs/log2.log testdata/test04/logs/log.log
./myRotate  testdata/archive testdata/test04/logs/log2.log
./myRotate testdata/test04/logs/log1.log
./myRotate testdata/test04/logs/log.log
*/
