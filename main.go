package main

import "fmt"

func main() {
	const message = "Hello, World!"
	fmt.Println(message)
	for i := 1;i<=5;i++{
		fmt.Println("*")
		for j:=1;j<i;j++{
			fmt.Print("*")
		}
	}
}
