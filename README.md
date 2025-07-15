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
```go
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

## Ways of declaring a variable
```go
name := "Leonardo"
var name = "Leonardo"
var name string = "Leonardo"
```

This is also possible:
```go
x1, x2 := 1, 2
var y1, y2 = 1, 2
```

## Constants
They need to be declared and assign a value in the same line.
```go
const myName string = "Leonardo"
```

## Functions
```go
func main() {
    printMe()
}

func printMe() {
}

func printMe2(printValue string) {
}
```

```go
func rValue1() int {
    return 1
}
```

Returning multiple values:

```go
func x() (int, int) {
    return 1, 1
}
```

## Ignoring a returning value
The _string_ value is ignored.
```go
someVal, _ := myFunc()
fmt.Println(someVal)

func myFunc() (int, string) {
	return -1, "Leo"
}

```

## Default value of _error_ 
The default value of the _error_ type is nil.

## Conditional switch statement
```go
switch remainder {
case 0:
    fmt.Println("The division was exact")
case 1, 2:              // 1 or 2
    fmt.Println("The division was close")
default:            // anything else
    fmt.Println("The division was not close")
}
```

## Arrays
- Fixed length
- Same Type
- Indexable
- Contiguous in memory

### How to initialize them
```go
var intArr [3]int32 = [3]int32{4, 2, 6}
intArr[1] = 123
```
or:
```go
intOther := [3]int32{4, 5, 6}
```
or inferring the size:
```go
intOther := [...]int32{5, 2, 65}
```
