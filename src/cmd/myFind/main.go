package main

import (
	"fileutils/internal/filefinder"
	"fileutils/internal/parser"
	"flag"
)

func main() {
	fileFlag, dirFlag, symlinkFlag, extFlag := parser.ParseFlagsFind()
	// Запускаем обход директории
	filefinder.WalkDirectory(flag.Args()[0], fileFlag, dirFlag, symlinkFlag, extFlag)
}
