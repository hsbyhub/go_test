package main

import "fmt"

type IAnimal interface {
	Speak() string
}

type Animal struct {
	Name string
	Age  string
}

type People struct {
	Animal
}

func (*People) Speak() string {
	return "Hei~"
}

type Sheep struct {
	Animal
}

func (*Sheep) Speak() string {
	return "Mie~"
}

func Speak(obj interface{}) {
	animal := obj.(IAnimal)
	fmt.Println(animal.Speak(), animal)
}

func main() {
	p := new(People)
	s := new(Sheep)
	Speak(p)
	Speak(s)
}
