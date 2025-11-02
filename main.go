package main

type Speck interface{
	Speck() string
}

type Dog struct {
	Name string
}

func (d Dog) Speck() string {
	return  "Woof!"
}
 type Person struct {
	Name string
 }

func (p Person) Speck() string {
	return "Hello!"
}

func makeSound(s Speck) {
	println(s.Speck())
}

func main() {
 dog := Dog{Name: "Buddy"}
 person := Person{Name: "Alice"}

 makeSound(dog)
 makeSound(person)
}