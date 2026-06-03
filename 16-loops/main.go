package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	println("\nFor like while loop")
	i := 1
	for i <= 10 { // for loop but like a while
		if i%2 == 0 {
			print(i, " ")
		}
		i++
	}
	println("\nunconditional loop")

	i = 0

	for {
		i++
		if i > 10 {
			break // breaks loop
		}

		if i%2 == 0 {
			continue // continue to the next iteration
		}
		print(i, " ")
	}

	println("\n no init only condition and postcondition")
	i = 1
	a, b := 0, 1
	for ; i <= 10; i += 1 {
		print(a, " ")
		a, b = b, a+b
	}

	println("\nmultiple init")
	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	println("nested loops")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			println("i:", i, "j:", j)
		}
	}

	println("break nested loops")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break
			}
			println("i:", i, "j:", j)
		}
	}

	println("break both loops")
	done := false

	for i := 1; i <= 5; i++ {
		if done {
			break
		}
		for j := 1; j <= 5; j++ {
			if j == 3 {
				done = true
				break
			}
			println("i:", i, "j:", j)
		}
	}

	println("break both loops using break outer")

outer: // label
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				if k == 3 {
					break outer
				}
				println("i:", i, "j:", j, "k:", k)
			}
		}
	}

exit:
	for {
		num := rand.IntN(999) // assume when ever the number is divisible by 17 break the loop
		switch {
		case num%2 == 0:
			println(num, "if even number")
			if num%17 == 0 {
				break exit
			}
		default:
			println(num, "is odd number")
			if num%17 == 0 {
				break exit
			}
		}
	}
}

// Go has only one loop that is for loop
