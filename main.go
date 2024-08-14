package main

import "fmt"

func main() {
	fmt.Println(foo())
}

func foo() string {
	return bar()
}

func bar() string {
	return "Hello World"
}
