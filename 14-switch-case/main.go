package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	day := uint8(rand.IntN(8))

	switch day {
	case 1:
		println("Sunday")
		// break no need to give explicitly break statement
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println(day, "invalid day")
	}

	char := 'G'

	switch char {
	case 'a', 'e', 'i', 0x6F, 0b01110101, 65, 0x45, 0b01001001, 'O', 'U':
		println(string(char), "is vovel")
	default:
		println(string(char), "is a consonent or a special character")
	}

	switch c := rand.IntN(255); c { // expression
	case 'a', 'e', 'i', 0x6F, 0b01110101, 65, 0x45, 0b01001001, 'O', 'U':
		println(string(c), "is vovel")
	default:
		println(string(c), "is a consonent or a special character")
	}

	num := rand.IntN(1000)

	switch { // empty switch if the case(s) has a condition
	case num%2 == 0:
		println(num, "is an even number")
	case num%2 != 0:
		println(num, "is an odd number")
	default:
		println(num, "default case")
	}

	// fallthrough
	// if fallthrough is used, it automtallically goes to the next case body , but not check then case condition
	// fallthough is ~= when the break is avoided in other programming languages
	println("\nright way of using fallthrough")

	num = 2
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	// false nagative of fallthrough
	// if the order of caszes are not properly given, dall through gives wrong results

	println("\nwrong way of using fallthrough")
	num = 2

	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	fmt.Println("Hello world")
	println("hello Wrold")
}

// switch case
