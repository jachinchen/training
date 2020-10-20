package main

import (
	"fmt"
)

//HelloWorld prints hello, username.
func HelloWorld(userName string) string {
	return fmt.Sprintf("Hi, %s ", userName)
}

func main() {
	a := 1
	fmt.Printf(HelloWorld("appleBoy"))
	fmt.Println(a, "一天就學會 Go 語言")

	fmt.Println("a >= 1")
}
