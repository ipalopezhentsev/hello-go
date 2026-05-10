package main

import (
	"fmt"
	"maps"
	"slices"
)

const (
	a int = 10
	b int = 20
	c     = "10"
)

func main() {
	// testPrimitives()
	//testSlices()
	// testSlicesOfSlices()
	// testStrings()
	testMaps()
	// testSets()
}

func testPrimitives() {
	fmt.Println(f())

	fmt.Printf("Speed of sound is %d m/s\n", 343)
	fmt.Printf(`Speed of sound is 
	!!!!DRUM ROLL!!!
	%d
`, 300_000)

	flag := false
	fmt.Printf("Flag is %t\n", flag)

	flag = true
	fmt.Printf("Flag is %t\n", flag)

	var a1 int8 = 1
	var a2 int16 = 32767
	var a3 int32 = int32(int16(a1) + a2)
	fmt.Println("a1+a2 =", a3)

	if a1 == 1 {
		fmt.Println("a1 is really 1!")
	}

	const x int32 = 10
	//	x++

	fmt.Println("package level a =", a)

	//this is array! it cannot grow
	var y = [...]int{1, 2, 3}
	//this is array too:
	// var y = [3]int{1, 2, 3}
	//this is NOT array! it's a slice! it can grow
	c := []int{1, 2, 3}
	fmt.Println("slice:", c)
	//go-like format:
	fmt.Printf("array: %#v\n", y)
	fmt.Printf("array: %v\n", y)
	var i = 2
	q1, q2 := fmt.Println("Out of bounds:", y[i])
	var q3 = 10
	fmt.Printf("%v %v %d\n", q1, q2, q3)

	// huge := []int{1000000: 1}
	// fmt.Printf("%v", huge)

	var qqq = [...]int{1, 2, 3}
	fmt.Println("Two arrays are equal? yes -", y == qqq)
	//doesn't compile! different type, slices can be compared only with null, without helpers
	//fmt.Println("Two arrays are equal?", c == qqq)

	t1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("matrix:", t1)
	fmt.Println("Len of array:", len(t1))
}

func testSlices() {
	//nil slice: (technically slice is always a struct {pointer to backing store, len, cap} so not nil pointer,
	//it's not a pointer (unlike map!), but in comparison it acts like real nil)
	var slice1 []int
	var slice2 []int = []int{1, 2}
	//doesn't compile
	//fmt.Printf("Two slices equal? %v", slice1 == slice2)
	fmt.Println("Slice1 is nil? yes -", slice1 == nil)
	fmt.Println("Slice2 is nil? no - ", slice2 == nil)
	//var iv = slice1[0]
	//yes
	//fmt.Printf("Panic? %v", iv)
	//only comparing of slices:
	fmt.Println("Equal via slices.Equal?", slices.Equal([]int{1, 2}, slice2))
	var nilSlice []int
	fmt.Println("Len of nil slice should be 0:", len(nilSlice))

	fmt.Println("Appending to nil slice:", append(nilSlice, 1))
	fmt.Println("Did this change source slice? no:", nilSlice)
	nilSlice = append(nilSlice, 1)
	fmt.Println("And now? yes:", nilSlice)
	//can we use append without assignment? no
	//append(nilSlice, 2)
	nilSlice = append(nilSlice, []int{3, 4, 5}...)
	fmt.Println("Appending another slice:", nilSlice)

	//make

	//both capacity=5 and length(!)=5, filled with zeroes, appending adds 6th element!
	s := make([]int, 5)
	s = append(s, 1)
	fmt.Println("len(s) is 6!:", len(s))
	//capacity 5 and length 0:
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1)
	fmt.Printf("s1 has len 1! %v, cap=%v\n", s1, cap(s1))
	//clear sets to 0 and keeps the len!
	clear(s1)
	fmt.Println("s1 still has len 1:", s1)

	//slice comparison (== allows only comparison with nil)

	//bizarre:
	var nilSlice1 []int
	var emptySlice []int = []int{}
	fmt.Printf("nilSlice1 0/0: len=%v, cap=%v\n", len(nilSlice1), cap(nilSlice1))
	fmt.Printf("emptySlice 0/0: len=%v, cap=%v\n", len(emptySlice), cap(emptySlice))
	//they are both 0/0 but are UNEQUAL:
	fmt.Println("nilSlice1 == emptySlice:", nil == emptySlice)
	//doesn't compile:
	// fmt.Println("nilSlice1 == emptySlice:", nilSlice1 == emptySlice)

	//copying:
	a := []int{1, 2, 3, 4}
	b := []int{}
	//shouldn't cause effect since b has zero length:
	copy(b, a[:2])
	b1 := make([]int, 2)
	copy(b1, a[:2])
	fmt.Println("b should still be empty:", b)
	fmt.Println("b1 should be [1,2]:", b1)

	//conversions

	//arr->slice - inherits capacity/backing storage
	arr := [...]int{1, 2, 3}
	slice := arr[:]
	slice[0] = 10
	fmt.Printf("arr is now [10 2 3]: %v, slice is too: %v\n", arr, slice)
	fmt.Printf("Capacity of slice should match length of array (%v): %v\n", len(arr), cap(slice))
	slice = append(slice, 20)
	//but now slices' buffer has been reallocated, it no longer modified parent array:
	slice[1] = 100
	fmt.Println("arr is still [10 2 3]:", arr)
	fmt.Println("slice is now [10 100 3 20]:", slice)

	//slice->arr - materializes slice to independent copy
	sl1 := []int{1, 2, 3}
	//copies first two elems from sl1
	ar1 := [2]int(sl1)
	//doesn't touch sl1:
	ar1[0] = 10
	fmt.Println("sl1 should be [1,2,3]:", sl1)
	fmt.Println("ar1 should be [10,2]:", ar1)

	//proof that slices are structs and are passed by value, so any addition
	//(even if within capacity!) has no effect outside!
	doOnArr := func(arr []int) {
		arr[0] = 100
		arr = append(arr, 200)
		//here it changed a !copy! of slice struct - even if there was capacity and it did not
		//reallocate backing store, it had to change the struct's len !field! - it will be invisible
		//to outside caller if we don't return the new value!
	}
	//cap=100
	a = make([]int, 0, 100)
	a = append(a, 1, 2, 3)
	//will pass a copy of a struct having three fields: pointer to storage, and len and cap ints
	doOnArr(a)
	//doOnArr will make only one change visible to us: change first element, because the copy's
	//storage pointer will initially be pointing at the same storage as ours:
	fmt.Println("After function doOnArr updates a, only first element changes to 100, 200 is not appended", a)

	//now repeat the same on pointers:
	doOnArrPtr := func(arr *[]int) {
		(*arr)[0] = 100
		//this will change memory to which pointer points, i.e. to original struct
		*arr = append(*arr, 200)
	}
	a = make([]int, 0, 100)
	a = append(a, 1, 2, 3)
	doOnArrPtr(&a)
	//so the outer function sees the append:
	fmt.Println("After function doOnArrPtr updates memory pointed by pointer, parent func will see the appended 200:", a)
}

func testSlicesOfSlices() {
	srcSlice := []int{1, 2, 3}
	subSlice := srcSlice[1:2]
	fmt.Println("Subslice is '[2]':", subSlice)
	//will modify parent buffer!
	subSlice[0] = 5
	fmt.Println("Subslice is '[5]':", subSlice)
	fmt.Println("Parent subslice is '[1,5,3]'", srcSlice)
	//appending to subslice will reuse capacity of parent slice, if there is some space!
	fmt.Printf("Subslice: len(1)=%v, cap(2)=%v\n", len(subSlice), cap(subSlice))
	subSlice = append(subSlice, 10)
	//this modified parent slice! because subSlice capacity was 2
	fmt.Println("Parent slice is now [1,5,10]:", srcSlice)
	fmt.Println("subSlice slice is now [5,10]:", subSlice)
	//subSlice is out of capacity of parent slice! so next append 'unties' it from parent slice!
	subSlice = append(subSlice, 11)
	fmt.Println("subSlice is now disconnected from parent slice:", subSlice)
	fmt.Println("srcSlice is disconnected from subSlice and is still [1,5,10]:", srcSlice)

	//full slice expressions:
	srcSlice2 := []int{1, 2, 3}
	//by last '2' we tell that index 2 of parent capacity is already the end of this slice
	subSlice1 := srcSlice2[1:2:2]
	//so this is gonna reallocate capacity on first append:
	subSlice1 = append(subSlice1, 5)
	subSlice1[0] = 100
	//so this has no effect on parent slice:
	fmt.Println("subSlice1 is [100,5]", subSlice1)
	fmt.Println("Parent slice is unchanged [1,2,3]:", srcSlice2)
}

func testStrings() {
	sRus := "й"
	sEng := "q"
	//strings store chars in utf-8
	//russian letter (code point) in utf-8 is two bytes
	//len returns in bytes, not code points

	//2:
	fmt.Println("len(sRus):", len(sRus))
	//1:
	fmt.Println("len(sRus):", len(sEng))

	//but subscript ("slice") returns bytes not code points!
	garbage := sRus[0]
	fmt.Printf("garbage and not %v: %v\n", sRus, string(garbage))

	sRusBytes := []byte(sRus)
	sRusRunes := []rune(sRus)
	fmt.Printf("sRusBytes=%v, sRusRunes=%v\n", sRusBytes, sRusRunes)

	//but for range iterates over runes! and i will show byte offset from beginning:
	for i, r := range "йц" {
		fmt.Println("At idx", i, "rune", r, "symbol", string(r))
	}
}

func testMaps() {
	m := map[string][]int{
		"key1":             {1, 2, 3},
		"key2":             {2, 3},
		"keyWithZeroValue": {},
		"keyWithNilValue":  nil,
	}
	fmt.Println("Map:", m)
	fmt.Println("m[key1] =", m["key1"])
	fmt.Println("m[missing] =", m["missing"])
	fmt.Println("m[missing]==nil? yes:", m["missing"] == nil)
	fmt.Println("m[keyWithZeroValue]==nil? no!:", m["keyWithZeroValue"] == nil)
	fmt.Println("m[keyWithNilValue]==nil? yes!:", m["keyWithNilValue"] == nil)

	//having append as a function allows it not to fail on nulls in java terms:
	q := append(m["missing"], 1)
	fmt.Println("Appended to missing key's value:", q)
	//but has it changed the map? no!
	fmt.Println("m[missing] again =", m["missing"])
	m["missing"] = q
	//and now yes:
	fmt.Println("m[missing] after direct update =", m["missing"])

	//values for missing key and for key with zero value are indisinguishable by simple indexing:
	nilValueForNonExistingKey := m["qqq"]
	var nilSlice []int
	fmt.Println("Missing equals zero val? yes:", slices.Equal(nilValueForNonExistingKey, []int{}))
	fmt.Println("Missing equals nilSlice? yes:", slices.Equal(nilValueForNonExistingKey, nilSlice))
	fmt.Println("Is nilValueForNonExistingKey nil? yes:", nilValueForNonExistingKey == nil)
	fmt.Println("Is nilSlice nil? yes:", nilSlice == nil)
	reallyExistingZeroValue := m["keyWithZeroValue"]
	//this is just go thing where []int (nil) is not same as []int{}
	fmt.Println("reallyExistingZeroValue same as nil? NO!:", reallyExistingZeroValue == nil)
	//but they are both empty:
	fmt.Printf("reallyExistingZeroValue slices.Equals to nil? YES!: %v\n", slices.Equal(reallyExistingZeroValue, nil))
	//(so with value type being a slice it does not fully allow to show this, but simple int would)
	//let's do quick demo with int map:
	{
		m := map[string]int{
			//this value will be indistinguishable from value of a missing key without comma ok idiom:
			"1": 0,
			"2": 2,
		}
		fmt.Println("m[1] is 0:", m["1"])
		fmt.Println("And m[100] is 0:", m["100"])
		existingZero, ok := m["1"]
		fmt.Println("But using comma idiom we can distinguish: existingZero =", existingZero, ", ok(true) =", ok)
		nonExistingZero, ok := m["100"]
		fmt.Println("nonExistingZero =", nonExistingZero, ", ok(false) =", ok)
	}

	//comma ok idiom allows to do 'contains key' semantic:
	nilValueForNonExistingKey, ok := m["qqq"]
	fmt.Printf("qqq key does not exist - yes: %v (value %v)\n", !ok, nilValueForNonExistingKey)
	//note that ":=" still works despite we already used "ok"!
	//because the rule is it's permitted if there are NEW variables
	//in the left part and we have a new var in addition to "ok"!
	zeroValueForExistingKey, ok := m["keyWithZeroValue"]
	//but we can do this as well:
	zeroValueForExistingKey, ok = m["keyWithZeroValue"]
	fmt.Printf("keyWithZeroValue key exists - yes: %v (value %v)\n", ok, zeroValueForExistingKey)

	//map doesn't have cap because it's not open addressing but bucket-based:
	//fmt.Printf("cap(m): %v, len(m): %v", cap(m), len(m))
	fmt.Println("len(m)=5:", len(m))

	//preallocating a map length - it will be 0 length but will have enough space to hold 1000 elems without rehashing:
	preallocatedMap := make(map[int]string, 1000)
	fmt.Println("preallocatedMap len(0) =", len(preallocatedMap))

	//maps are not comparable without helpers:
	m1 := map[int]string{1: "1", 2: "2"}
	m2 := map[int]string{1: "1", 2: "2"}
	//doesn't compile:
	// fmt.Printf("m1==m2: %v", m1 == m2)
	fmt.Println("m1 maps.Equal m2 - YES:", maps.Equal(m1, m2))

	//deletion:
	delete(m1, 1)
	fmt.Println("m1 is now {2: 2}:", m1)

	//clear clears all:
	clear(m2)
	fmt.Println("m2 is empty:", m2)

	//unlike slices, maps are not structs but pointers and therefore allow adding inside function:
	doOnMap := func(m map[string]int) {
		m["1"] = 100
		m["100"] = 100
	}
	mm := map[string]int{
		"1": 1,
		"2": 2,
	}
	//will pass a copy of !pointer! so everything done inside the func will be visible outside
	doOnMap(mm)
	fmt.Println("mm modified by function, should be [1:100, 2:2, 100:100]", mm)
}

func testSets() {
	//set via dummy 'true' bool value - takes extra byte for bool but easy to use:
	s := map[int]bool{1: true, 100: true}
	fmt.Println("s has 1 and 100:", s[1] && s[100])
	fmt.Println("s doesn't have 2:", !s[2])

	//set via values of type "empty struct" - doesn't take space for vals but clumsier to use:
	// s1 := map[int]struct{}{1: struct{}{}, 100: struct{}{}}
	//shorter form:
	s1 := map[int]struct{}{1: {}, 100: {}}
	_, ok := s1[1]
	fmt.Println("s1 has 1:", ok)
	_, ok = s1[2]
	fmt.Println("s1 doesn't have 2:", !ok)
}
