package main

import "fmt"

type Human struct {
	Age int
}

func (h Human) getAge() {
	fmt.Println(h.Age)
}

type Action struct {
	AgeGetter
}

type AgeGetter interface {
	getAge()
}

func main() {
	action := Action{Human{30}}
	action.getAge()
}
