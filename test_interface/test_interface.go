package main

import "fmt"

type Foo struct {
	num int
}

func Assign100(v interface{}) {
	v.(*Foo).num = 100

}

func main() {
	foo := Foo{
		num: 0,
	}
	Assign100(&foo)
	fmt.Println(foo)
}
