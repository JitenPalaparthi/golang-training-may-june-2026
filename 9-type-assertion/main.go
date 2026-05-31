package main

func main() {

	//num1, num2, num3, float1, float2 := 9444, uint64(3432423), uint16(43242), float32(676.45), float64(45345.345)

	//var any1 interface{} = 24234234 // int value
	var any1 any = 24234234 // int value

	// any1 is data and type pointerns , data pointer is address to the value 24234234, type pointer is an address to the type int

	//var num int = int(any1) // With any type castint, or conversion does not work, it works on type assertion
	var num1 int = any1.(int) // With any type castint, or conversion does not work, it works on type assertion

	//any.(Type)
	println(num1)

	any1 = 64546.456456

	// var num2 = any1.(int) // it failed to assert bcz any1 has the type of float64 but trying to assert to int which is invalid leads to panic/runtime error
	// println(num2)

	num2, done := any1.(int) // it failed to assert bcz any1 has the type of float64 but trying to assert to int which is invalid leads to panic/runtime error
	if done {
		println("successfully asserted to int", num2)
	} else {
		num2, done := any1.(float64)
		if done {
			println("successfully asserted to float64", num2)
		} else {
			println("failed to assert")
		}
	}

	println("hello world")

}
