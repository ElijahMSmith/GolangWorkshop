package main

import "fmt"

func testBasicStructures() {
	array := create5IntArray(10)
	for index, element := range array {
		fmt.Printf("Index %d: %d\n", index, element)
	}
	fmt.Println()

	slice, err := createIntSlice(-1, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice)
	}

	slice, err = createIntSlice(3, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice, len(slice))
	}

	slice = append(slice, 10)
	newLen := len(slice)
	fmt.Println(slice, newLen)
	fmt.Println()

	newMap := make(map[string]int)
	newMap["one"] = 1
	newMap["two"] = 2

	for key, val := range newMap {
		fmt.Printf("%s: %d\n", key, val)
	}

	val, ok := newMap["two"]
	fmt.Println(val, ok)
	val, ok = newMap["three"]
	fmt.Println(val, ok)
}

func create5IntArray(defaultVal int) [5]int {
	// Array is a value type, so cannot be nil
	// Gets copied when passing into a function, which can be really slow
	// Fixed length, and the length we provide can't be a variable
	created := [5]int{}
	for i := 0; i < 5; i++ {
		created[i] = defaultVal
	}
	return created
}

func createIntSlice(length int, defaultVal int) ([]int, error) {
	// Slice is a reference type, which means we don't copy it when sending it into a function
	// We can return this as nil since it is a reference
	// We can also make this whatever variable length we want, and it grows/shrinks as we want!
	if length < 0 {
		return nil, fmt.Errorf("Provided length %d cannot be negative!", defaultVal)
	}

	created := make([]int, length)
	for i := 0; i < length; i++ {
		created[i] = defaultVal
	}
	return created, nil
}
