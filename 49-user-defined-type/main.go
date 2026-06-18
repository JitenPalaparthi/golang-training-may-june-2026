package main

import "fmt"

func main() {

	var (
		num  int    = 100
		num1 MyInt1 = 200
		num2 MyInt2 = 300
		num3 MyInt3 = 400
		num4 MyInt4 = 500

		float1 float32 = 87.54
	)

	s := MyInt1(num).ToString()
	println("int Tostring:", s)

	sq := MyInt2(num).Sq()
	println("int Sq:", sq)

	cb := MyInt3(num).Cube()
	println("int Cube:", cb)

	//

	s1 := num1.ToString()
	println("num1 Tostring:", s1)

	sq1 := MyInt2(num1).Sq()
	println("num1 Sq:", sq1)

	cb1 := MyInt3(num1).Cube()
	println("num1 Cube:", cb1)

	//
	s2 := MyInt1(num2).ToString()
	println("num2 Tostring:", s2)

	sq2 := num2.Sq()
	println("num2 Sq:", sq2)

	cb2 := MyInt3(num2).Cube()
	println("num2 Cube:", cb2)

	s3 := MyInt1(num3).ToString()
	println("num3 Tostring:", s3)

	sq3 := MyInt2(num3).Sq()
	println("num3 Sq:", sq3)

	cb3 := num3.Cube()
	println("num3 Cube:", cb3)

	s4 := MyInt1(num4).ToString()
	println("num4 Tostring:", s4)

	sq4 := MyInt2(num4).Sq()
	println("num4 Sq:", sq4)

	cb4 := MyInt3(num4).Cube()
	println("num4 Cube:", cb4)

	s5 := MyInt1(float1).ToString()
	println("float1 Tostring:", s5)

	sq5 := MyInt2(float1).Sq()
	println("float1 Sq:", sq5)

	cb5 := MyInt3(float1).Cube()
	println("float1 Cube:", cb5)

	var float2 = 34.56
	println(int(float2))
}

type MyInt1 int

type MyInt2 int

type MyInt3 int

type MyInt4 MyInt3

//type Integer = int // Alias not creating a type

func (m MyInt1) ToString() string {
	return fmt.Sprint(m)
}

func (m MyInt2) Sq() int {
	return int(m * m)
}

func (m MyInt3) Cube() int {
	return MyInt2(m).Sq() * int(m) // return int(m * m * m)
}
