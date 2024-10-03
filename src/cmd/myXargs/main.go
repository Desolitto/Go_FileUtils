package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myXargs <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[1:]

	var args []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			args = append(args, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "stdin: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(command[0], append(command[1:], args...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

/*
go build -o myXargs ./cmd/myXargs
go build -o myFind ./cmd/myFind
go build -o myWc ./cmd/myWc
echo -e "testdata/test03/a\ntestdata/test03/b\ntestdata/test03/c" | ./myXargs ls -la
*/
