package main

import "fmt"

func testBasics() {
	fmt.Println("Hello, World!")

	var i int = 4
	fmt.Printf("i is %d\n", i)

	j := 8
	fmt.Println(i - j)

	multiply(i, j)
	x, y, z, err := getThreeNumbers()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The product is %d\n", multiply(x, y)*z)
	}

	loops()
}

func multiply(i int, j int) int {
	return i * j
}

func getThreeNumbers() (int, int, int, error) {
	var x, y, z int
	fmt.Println("Please input three numbers:")
	accepted, err := fmt.Scan(&x, &y, &z)

	fmt.Println(accepted)
	fmt.Println(err)

	return x, y, z, err
}

func loops() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	j := 0
	for j < 10 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	k := 0
	stop := 10
	for {
		fmt.Printf("%d ", k)
		k++

		if k >= stop {
			break
		}
	}

	// For range loop later
}
