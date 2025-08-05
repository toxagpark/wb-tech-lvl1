package main

import "fmt"

type human struct {
	name        string
	blood_group int
}

func (h human) sayGroup() {
	fmt.Printf("My group: %d\nMy name: %s\n", h.blood_group, h.name)
}

type action struct {
	human
	operation string
}

func (a action) doOperation() {
	fmt.Printf("My operation: %s \n", a.operation)
}

func main() {
	action := action{
		human: human{
			name:        "Tim",
			blood_group: 2,
		},
		operation: "transfusion",
	}
	action.doOperation()
	action.sayGroup()
}
