package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		return "Hello, world!"
	}
	return "Hello, " + name + "!"
}
