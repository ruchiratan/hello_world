package main

import "fmt"

func main() {
	fmt.Println("hello Ruchi!")
	fmt.Println("Golang is very simple")
}

func init() {
	fmt.Println("I run before main")
}