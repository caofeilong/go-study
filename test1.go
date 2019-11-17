package main

import "fmt"

func j() A {
	return e()
}

type A struct {
	name string
}

func e() A {
	return A{name: "123"}
}

func y() (string, string) {
	return "a", "b"
}

func main() {
	var b string
	var a string

	if 1 == 1 {
		a, b = y()
		if a == "a" {
			fmt.Println("--a--")
		}
	}

	fmt.Println(b)
}
