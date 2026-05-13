package main

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  uint8
}

type Person1 struct {
	name string
	age  uint8
}

func TestStructs(t *testing.T) {
	p := Person{
		name: "Ilya",
		age:  42,
	}
	fmt.Println("p is", p)
	p1 := Person1{
		name: "Ilya",
		age:  42,
	}
	fmt.Println("p1 is", p1)

	//doesn't compile:
	// if (p != p1) {
	// 	t.Errorf("")
	// }
	//but this one should:
	var p2 Person = Person(p1)
	if p != p2 {
		t.Error()
	}

	//but can be compared with values of its own type:
	pNastya := Person{
		name: "Nastya",
		age:  40,
	}
	if p == pNastya {
		t.Error()
	}
	pIlya2 := Person{
		name: "Ilya",
		age:  42,
	}
	if p != pIlya2 {
		t.Error()
	}

	//structs are not pointers, nil is not allowed...
	//but it's zero value struct having all fields set to zero val:
	var zeroPerson Person
	//doesn't compile:
	// if nilPerson != nil {
	// 	t.Error()
	// }
	if zeroPerson.name != "" {
		t.Error()
	}
	if zeroPerson.age != 0 {
		t.Error()
	}
}

func TestAnonymousStructs(t *testing.T) {
	s := struct {
		name string
		age  uint8
	}{
		"Ilya",
		42,
	}
	//unlike all non-anonymous structs, anonymous struct can be compared
	//to another anonymous or non-anonymous struct!
	if (s != Person{name: "Ilya", age: 42}) {
		t.Error()
	}
	if s.age != 42 || s.name != "Ilya" {
		t.Error()
	}
}
