package main

import "fmt"

// for, break, continue, and goto

func main() {
	// execution1()
	// execution2()
	execution3()
}

func execution1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("Infinite loop")
	// }

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	for i := 0; i < 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println(i)
	}


	for i := 0; i < 5; i++ {

		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

}


func execution2() {
	// Go lets you give a loop a label

	outer:
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			if i == 1 && j == 1 {
				break outer
				// breaks the outer loop, not just the inner one.
			}

			fmt.Println(i, j)
		}
	}


	outer2:
	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			if j == 1 {
				continue outer2
				// continue outer skips to the next iteration of the outer loop.
			}

			fmt.Println(i, j)
		}
	}

}


func execution3() {
	i := 0

start:
    fmt.Println(i)

    i++

    if i < 5 {
        goto start // It jumps directly to start
    }
}