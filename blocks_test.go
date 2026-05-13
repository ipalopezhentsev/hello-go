package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestShadowing(t *testing.T) {
	x := 5
	if x == 5 {
		//!!shadowing
		x := 6
		if x != 6 {
			t.Error()
		}
	}
	//original x didn't change!
	if x != 5 {
		t.Error()
	}
}

func TestShadowingUniverseBlock(t *testing.T) {
	t.SkipNow()
	//if you erase SkipNow - the below will compile and go test will go fine!
	//just go vet won't work...
	false := true
	nil := 10
	if false != true {
		t.Error()
	}
	if nil != 10 {
		t.Error()
	}
}

func TestIf(t *testing.T) {
	//n will be scoped to just if body
	if n := rand.Intn(10); n > 5 {
		fmt.Println("good")
	}
	//n undefined:
	//fmt.Println(n)
}

func TestFor(t *testing.T) {
	//simple for:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//init before:
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//increment inside:
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	//condition only (while):
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	if i != 10 {
		t.Error()
	}

	//do-while (do-until actually)
	i = 0
	for {
		fmt.Println(i)
		if !(i < 10) {
			break
		}
		i++
	}
	if i != 10 {
		t.Error()
	}

	//rangeint:
	for i := range 10 {
		fmt.Println(i)
	}

	//range for slices:
	slice := []int{10, 20, 30, 40}
	for i, val := range slice {
		if val != (i+1)*10 {
			t.Error()
		}
	}

	//range for maps:
	m := map[string]int{
		"1": 1,
		"2": 2,
	}
	//just keys:
	var s string = ""
	for k := range m {
		s += k
	}
	//order is random!
	if s != "12" && s != "21" {
		t.Errorf("s=%v", s)
	}
	for range 5 {
		//different iterations will output in random order!
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	//range over strings moves over runes, not bytes!
	for i, r := range "ый" {
		fmt.Println("At byte offset ", i, "there is rune", string(r))
	}
}

func TestForRangeGivesACopyOfValue(t *testing.T) {
	//subscript of slice returns not a copy but actual value, even for structs!

	slOfInt := []int{1, 2}
	slOfInt[0]++
	if slOfInt[0] != 2 {
		t.Error()
	}

	type A struct {
		a int
	}
	slOfA := []A{{a: 1}, {a: 2}}
	//subscript will NOT return a copy (unlike maps!)
	slOfA[0].a++
	if slOfA[0].a != 2 {
		t.Error()
	}

	//mapOfA := map[string]A {"1": A{a:1},"2": A{a:2}}
	//doesn't compile:
	//mapOfA["1"].a++

	//and now the hero! for-range val is a copy, unlike direct subscript!
	slOfA = []A{{a: 1}, {a: 2}}
	for _, v := range slOfA {
		//goes nowhere...
		v.a++
	}
	if slOfA[0].a != 1 || slOfA[1].a != 2 {
		t.Error()
	}
}

func TestSwitch(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	//shortuct for traversing only keys
	i := 0
	for k := range m {
		switch k {
		case "1":
			i++
		case "2":
			i++
		default:
			t.Error()
		}
	}
	if i != 2 {
		t.Error()
	}
}

func TestBlankSwitch(t *testing.T) {
	nums := []int{2, 3, 4}
	numOds := 0
	numEvens := 0
	for _, v := range nums {
		//ugly but to show blank switch (";" is required!)
		// switch q := v; {
		switch {
		case v%2 == 0:
			numEvens++
		default:
			numOds++
		}
	}
	if numOds != 1 || numEvens != 2 {
		t.Errorf("numOds: %v, numEvens: %v", numOds, numEvens)
	}
}
