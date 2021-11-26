package main

import "fmt"

func main() {
	var a int
	var b int
	var ch string

	fmt.Println("Enter the a ")
	fmt.Scanln(&a)
	fmt.Println("Enter the b ")
	fmt.Scanln(&b)
	fmt.Println("Enter the operations:(+,_,*,/)")
	fmt.Scanln(&ch)
	if ch == "+" {
		fmt.Println("Sum is:\n", a+b)
	} else if ch == "-" {
		fmt.Println("Sub is :\n", a-b)
	} else if ch == "*" {
		fmt.Println("Multi is: \n", a*b)
	} else if ch == "/" {
		fmt.Println("Div is: \n", a/b)
	} else {
		fmt.Println("Invalid Operation")
	}

}
