package main

// escape analysis -> what gets escaped to heap memory
// go build -gcflags="-m" .
func main() {
	var num1 int = 100
	Incr(num1)

	println("Value of num1:", num1)

	ptr1 := new(num1)

	//IncrR(&num1)

	IncrR(ptr1)

	println("Value of num1:", num1)

	ptr2 := IncrP(num1)
	println("Value of num1:", num1)
	println("Value of ptr2:", *ptr2) // gives the pointer which is created inside the func

	ptr3 := new(num1)

	IncrInP(num1, ptr3)

	println("Value of ptr3", *ptr3)

	arr1 := [10]int{} // in stack most probably
	println(&arr1)

	arr2 := [100000]int{} // in heap most probably
	println(&arr2)

	slice1 := [10]int{} // in stack most probably
	println("address of slice1", &slice1, "original data of slice1", &slice1[0])

	slice2 := [100000]int{} // in heap most probably
	println(&slice2)

	slice3 := []int{}
	//fmt.Println(slice3) // never use fmt.Println in production, why it escapes all valies to heap which are used in fmt.Println
	println(&slice3)

	for i := range 10000 {
		slice3 = append(slice3, i)
	}
	println("address of slice3", &slice3, "original data of slice3", &slice3[0])
}

// 0x292dfb48ef8
// 0x292dfb42000

func Incr(n int) { // call or pass by value
	n++
}

// call by ref
func IncrR(p *int) { // call or pass by ref
	*p++
}

// return a pointer
// This is considered to be a very big problem called dangling pointer
// w.r.t c or c++ , or rust (rust does not compile it at all)
// go handles differently
func IncrP(n int) *int {
	ptr := new(int) // var ptr *int
	*ptr = n + 1
	return ptr
}

// not a dangling, bcz the out is coming from the called not created inside the function
func IncrInP(n int, out *int) {
	*out = n + 1
}
