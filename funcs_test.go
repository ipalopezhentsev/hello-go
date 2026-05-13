package main

import (
	"errors"
	"testing"
)

func TestFunctionWithMultipleReturnVals(t *testing.T) {
	q := func(slice []int) (capacity int, length int) {
		return cap(slice), len(slice)
	}
	c, l := q(make([]int, 3, 10))
	if c != 10 || l != 3 {
		t.Error()
	}

	//surprisingly, we can ignore all results of multi-return value func,
	//but not some of them:
	//permitted:
	q([]int{1, 2})

	//not permitted:
	//c = q([]int {1,2})

	//named return arguments are actually defined inside func's body:
	//(but it's a bad practice)
	q1 := func(slice []int) (capacity int, length int) {
		//so we should be careful not to shadow them...
		capacity = cap(slice)
		length = len(slice)
		//!!!
		return
	}
	if c, l := q1([]int{0, 1}); c != 2 || l != 2 {
		t.Error()
	}
}

func TestFunctionWithErrors(t *testing.T) {
	div := func(a int, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("Cannot divide by 0")
		}
		return a / b, nil
	}

	if a, err := div(8, 2); err != nil {
		t.Error()
	} else {
		if a != 4 {
			t.Errorf("a=%v", a)
		}
	}

	if _, err := div(8, 0); err == nil {
		t.Error()
	} else {
		if err.Error() != "Cannot divide by 0" {
			t.Error()
		}
	}
}

func TestFunctionsAsValues(t *testing.T) {
	greeter := func(name string, helloLocalizer func() string) string {
		return helloLocalizer() + name
	}
	if greeter("Ilya", func() string { return "Hello " }) != "Hello Ilya" {
		t.Error()
	}
	if greeter("Илья", func() string { return "Привет, " }) != "Привет, Илья" {
		t.Error()
	}

	//we can introduce type name for the function.
	//(but we still have to use full declaration in lambda...)
	type helloLocalizer func() string
	greeter1 := func(name string, localizer helloLocalizer) string {
		return localizer() + name
	}
	if greeter1("Ilya", func() string { return "Hello " }) != "Hello Ilya" {
		t.Error()
	}
}

func TestClosures(t *testing.T) {
	a := 10
	f := func() {
		if a != 10 {
			t.Error()
		}
		a = 20
	}
	f()
	//!!so closure modified our on stack value!
	//(also read comment at the end, in general it's not always on stack)
	if a != 20 {
		t.Error()
	}

	//it also sees new modified value:
	//f()

	//structs would have the same behavior!
	//note that it's ok to return such closure from a function!
	//in this case compiler will detect that captured variables "escape the stack"
	//and it will convert them to heap and pointers!
}
