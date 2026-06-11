package main

import "fmt"

func main() {

	var num1 int = 100

	var ptr1 *int // ptr1 itself has its own address

	// if there is a pointer , you dont give any value , it would be nil

	if ptr1 == nil {
		println("ptr1 is nil")
	}

	ptr1 = &num1

	fmt.Println("Value of num1:", num1)
	fmt.Println("Address of num1:", &num1)
	fmt.Println("Value of num1 thru pointer:", *ptr1)
	fmt.Println("Address of num1 , stored in ptr1:", ptr1)
	fmt.Println("Address of ptr1:", &ptr1)

	// any variable that holds the address, can be dereferenced, as long as there is address

	*(&num1) = 200 // is equal to num1= 200
	fmt.Println("Value of num1:", num1)

	*ptr1 = 300 // the pointer contains the address of num1
	fmt.Println("Value of num1:", num1)

	*(*(&ptr1)) = 400
	fmt.Println("Value of num1:", num1)

	var ptr2 **int = &ptr1

	**ptr2 = 500
	fmt.Println("Value of num1:", num1)

	//var ptr2 *int = &100 // not possible in Go is possible in rust

	// ptr3 := new(int)
	// *ptr3 = 100

	var ptr3 *int = new(100) // new allocates memory and returns a pointer
	*ptr3 = 1200

	num1 = 600

	Incr(num1)

	println("Value of num1:", num1)

	IncrR(&num1)

	println("Value of num1:", num1)

	IncrR(ptr1) // Ptr1 is a pointer that holds the addree of num1
	println("Value of num1:", num1)

	// *ptr1 --> 602

	Incr(*ptr1) // *ptr1 is a value
	println("Value of num1:", num1)

	arr1 := [2]int{10, 20}
	var ptr_arr1 = &arr1

	fmt.Printf("address of arr1:%p\n", &arr1)
	fmt.Printf("address of arr1 thru pointer:%p\n", ptr_arr1)

	//for _, v := range *ptr_arr1 {
	for _, v := range ptr_arr1 { // go can understand without defref
		print(v, ">>")
	}
	println()
	(*ptr_arr1)[0] = 1111 // This is ideally but
	ptr_arr1[1] = 2222    // can do this with out any issues

	for _, v := range ptr_arr1 { // go can understand without defref
		print(v, ">>")
	}
	println()

	var arr_ptr1 [3]*int // array holds the pointers
	arr_ptr1[0], arr_ptr1[1] = new(100), new(200)

	for _, p := range arr_ptr1 {
		if p != nil {
			print(*p, " >> ")
		}
	}
	println()
}

func Incr(n int) { // call or pass by value
	n++
}

func IncrR(p *int) { // call or pass by ref
	*p++
}
