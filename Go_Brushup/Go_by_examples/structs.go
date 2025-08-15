package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name, age}
}

func (p *Person) Bio() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	personObj := newPerson("Adam", 1)
	fmt.Println(personObj)
	fmt.Println(personObj.Name)
	fmt.Println(personObj.Age)

	fmt.Println(personObj.Bio())
}
