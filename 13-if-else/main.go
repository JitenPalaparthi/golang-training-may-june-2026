package main

import "strconv"

func main() {
	a, b := 10, 20

	println(a+b+10+(a*a+b*b)+(a+b)/2+(a+b%2)-100+250/2*2 >= 720 || false && true)

	if a+b+10+(a*a+b*b)+(a+b)/2+(a+b%2)-100+250/2*2 >= 720 || false && true {
		println("yes,true")
	} else {
		println("no, false")
	}

	str1 := "56354353"

	if v, err := strconv.Atoi(str1); err != nil {
		println(err.Error())
	} else {
		println("successfully converted", v)
	}

	//println(v, err.Error())

	age, gender := 42, 'm'

	if age >= 18 && (gender == 'f' || gender == 0x46) {
		println("She is eligible for marriage, because age is ", age)
	} else if age >= 21 && (gender == 0b01101101 || gender == 'M') { // 0b01101101 ==  0x6D == 109 == 'm'
		println("He is eligible for marriage, because age is ", age)
	} else {
		println("not eligible for marriage")
	}

}
