package main

var (
	Global int = 100
)

func main() {

	str1 := "Hello World"
	str2 := "Hello, 世界"

	for i := range 10 { // not about the value of i , but about the numnber iterations here they are 10
		println(i + 1)
	}

	for i, v := range str1 { // i is index, v is value in byte
		println(i, "-->", string(v))
	}

	println("accessing each byte of a str using normal loop")

	for i := 0; i < len(str2); i++ {
		println(i, "-->", string(str2[i]))
	}

	for i, v := range str2 { // i is index, v is value in byte
		println(i, "-->", string(v))
	}

	str2 = "Hello, 世界"

	for _, char := range str2 {
		print(string(char), " >> ")
	}
	println()

	//for i, _ := range str2 {
	for index := range str2 {
		print("index: ", index, " >> ")
	}
	println()

	println(Global, &Global)

}

// range loop is used for range of numbers, strings,arrays, slices, maps, channels
// what range loop returns is different based on the construct
