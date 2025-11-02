package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func giveRaise(e *Employee, giveRaise int) {
	e.Salary += giveRaise
}

func main() {
	emp := Employee{Name: "John", Salary: 50000}
	giveRaise(&emp, 5000)
	fmt.Println("After raise: ", emp)
}
