package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Eat()
}

type Dog struct {
	name string
}

func (dog Dog) Eat() {
	fmt.Println(dog.name + " is eatting...")
}

func main() {
	d := Dog{name: "hasik"}
	dp := &Dog{name: "hasik"}
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(dp))
}
