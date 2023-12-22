package main

import "fmt"

func naming(name string) string {
	return "Hi " + name
}

func main() {
	var name string
	fmt.Println("Your name")
	fmt.Scanln(&name)
	fmt.Println(naming(name))
}
