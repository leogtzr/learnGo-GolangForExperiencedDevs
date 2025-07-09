# learnGo-GolangForExperiencedDevs
Code Snippets for a Go course: Golang for Experienced Devs

## Six Main Points about Go
### 1 - Statically Typed Language
```Go
var myVariable string
// or
var myVariable = "myString"
```

### 2 - Strongly Typed Language
In Go, we cannot do something like this:
```Go
a = 1
b = "two"
c = a + b
```

### 3 - Go is compiled
### 4 - Fast Compile Time
### 5 - Built In Concurrency
Goroutines are great.

### 6 - Simplicity


## Starting a new project

```bash
go mod init <project-name>
```

### A Hello World

$ tree
.
├── cmd
│   └── tutorial1
│       └── main.go
└── go.mod

2 directories, 2 files
$ cat cmd/tutorial1/main.go
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

Building:
```bash
go build cmd/tutorial1/main.go
```

Running:
```bash
./main
```
or:
```bash
go run cmd/tutorial1/main.go
```

## Variable Types
```go
var intNum int8 // or int16, int32, int64
```
```go
var intNum uint // or uint8, uint16, uint64 
```
```go
var floatNum float64 // or float32
```

### We cannot perform operations with mixed types
The following leads to a compile error:
```go
var floatNum32 float32 = 10.1
var intNum int32 = 2
var result float32 = floatNum32 + intNum
```

For multiline string variables
var myName string = `
Leonardo
Gutierrez
Ramirez
`

To handle fancy strings, it is better to use a specialized package such as: _unicode/utf8_, and:
```go
fmt.Println(utf8.RuneCountInString("Y"))
```

## Default values for types
uint, uint8, uint16, uint32, uint64
int, int8, int16, int32, int64
float32, float64
rune
is 0 ...
