# Go FileUtils

## Overview

Go FileUtils is a command-line utility designed for performing various file system operations. This tool allows users to find files, count them, execute commands, and archive log files. It is developed using the Go programming language and provides a user-friendly interface for working with files and directories.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Getting Started](#getting-started)
4. [Usage](#usage)
    1. [Exercise 00: Finding Files](#exercise-00-finding-files)
    2. [Exercise 01: Counting Files](#exercise-01-counting-files)
    3. [Exercise 02: Running Commands](#exercise-02-running-commands)
    4. [Exercise 03: Archiving Logs](#exercise-03-archiving-logs)
5. [Project Structure](#project-structure)

## Introduction

Working with file systems is one of the core tasks in programming. The Go FileUtils project provides a set of utilities that simplify tasks such as finding files, counting them, executing commands, and archiving log files. This tool can be beneficial for both developers and system administrators.

## Features

- Find files and directories using various criteria.
- Count lines, words, and characters in text files.
- Execute commands using input from standard input.
- Implement log rotation to archive old log files, preventing indefinite growth of single log files.

## Getting Started

### Prerequisites

- Go programming language installed (version 1.16 or higher recommended)
- Git for version control

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/Go_FileUtils.git
    ```
2. Navigate to the project directory:
    ```bash
    cd Go_FileUtils
    ```
3. Build the application:
    ```bash
    go build -o myFind ./src/ex00/myFind.go
    go build -o myWc ./src/ex01/myWc.go
    go build -o myXargs ./src/ex02/myXargs.go
    go build -o myRotate ./src/ex03/myRotate.go
    ```

## Usage

### Exercise 00: Finding Files

To find files and directories, use the `myFind` command:

```bash
# Find all files, directories, and symbolic links in /foo
./myFind /foo

# Find only files with the '.go' extension
./myFind -f -ext 'go' /path/to/dir

# Find directories only
./myFind -d /path/to/dir
```
or use Makefile:
```bash
make myFind
```

### Exercise 01: Counting Files

To count lines, words, or characters in text files, use the `myWc` command:

```bash
# Count words in input.txt
./myWc -w input.txt

# Count lines in multiple files
./myWc -l input2.txt input3.txt

# Count characters in files
./myWc -m input4.txt input5.txt input6.txt
```
or use Makefile:
```bash
make myWc
```

### Exercise 02: Running Commands

To execute commands using input from standard input, use `myXargs`:

```bash
# Executes 'ls -la' on /a, /b, /c
echo -e "/a\n/b\n/c" | ./myXargs ls -la

# Find '.log' files and count lines
./myFind -f -ext 'log' /path/to/logs | ./myXargs ./myWc -l
```
or use Makefile:
```bash
make myXargs
```

### Exercise 03: Archiving Logs

To archive log files, use the `myRotate` command:

```bash
# Archive a single log file
./myRotate /dataForTest/dataForMyRotate/app1.log

# Archive multiple log files into a specific directory
./myRotate -a /dataForTest/ /dataForTest/dataForMyRotate/app1.log /dataForTest/dataForMyRotate/app2.log
```
or use Makefile:
```bash
make myRotate
```

## Project Structure

```plaintext
Go_FileUtils/
│
├── src/
│   ├── cmd/
│   │   ├── myFind/
│   │   │   └── main.go          # Main command for finding files
│   │   ├── myWc/
│   │   │   └── main.go          # Main command for counting files
│   │   ├── myXargs/
│   │   │   └── main.go          # Main command for executing commands
│   │   └── myRotate/
│   │       └── main.go          # Main command for archiving logs
│   │  
│   ├── internal/
│   │   ├── filefinder/
│   │   │   └── finder.go        # Logic for finding files and directories
│   │   ├── parser/
│   │   │   └── parser.go        # Logic for parsing command-line arguments
│   │   ├── rotate/
│   │   │   └── rotate.go        # Logic for archiving log files
│   │   └── wc/
│   │       └── wc.go            # Logic for counting lines, words, and characters
│   │  
│   ├── testdata/
│   │   ├── archive/             # Directory for testing rotate
│   │   ├── test01/              # Test data for various functionalities
│   │   ├── test02/
│   │   ├── test03/
│   │   └── test04/
│   │ 
│   ├── go.mod                   # Go module file that defines dependencies and module information.
│   └── Makefile                 # Commands for building and running the project
├── .gitignore                   # Specifies files and directories to be ignored by Git
└── README.md                    # Project documentation

```
### Explanation of the Directories

- **`go.mod`**: Go module file that defines dependencies and module information.
- **`src/`**: Contains the source code for the project, organized into subdirectories for each component.
- **`cmd/`**: Contains the main commands for the utilities:
  - **`myFind/`**: Contains the main command for finding files.
  - **`myWc/`**: Contains the main command for counting files.
  - **`myXargs/`**: Contains the main command for executing commands.
  - **`myRotate/`**: Contains the main command for archiving logs.
- **`internal/`**: Contains internal packages that implement core functionality:
  - **`filefinder/`**: Logic for finding files and directories.
  - **`parser/`**: Logic for parsing command-line arguments.
  - **`rotate/`**: Logic for archiving log files.
  - **`wc/`**: Logic for counting lines, words, and characters in files.
- **`testdata/`**: Contains sample data for testing various functionalities, organized into subdirectories for different test scenarios.
- **`Makefile`**: Contains commands for building and running the project.
- **`.gitignore`**: Specifies files and directories to be ignored by Git.
- **`README.md`**: Provides documentation for the project, including setup instructions and usage guidelines.
