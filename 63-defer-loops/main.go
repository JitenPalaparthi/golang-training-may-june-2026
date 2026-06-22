package main

func main() {

	checkDefer() // normally executed w.r.t main
	a := 10

	defer func(a int) {
		println("defer value of global a using call by value:", a)
	}(a)

	defer func() {
		println("defer value of global a:", a)
	}()

	defer func(a *int) {
		println("defer value of global a using call by ref:", *a)
	}(&a)

	a += 10
	println("a in normal execution:", a)

	defer println("end of main")

	//func() { // uncomment and check
	for i := range 10 {
		defer func() {
			println("i:", i, "sq:", i*i)
		}()
		println("Normal I:", i)
	}
	//}()
	defer println("mid of main")

	println("start of main")

	str := "Hello World!"

	for _, c := range str {
		defer print(string(c))
	}

}

func checkDefer() {
	defer println("end of checkDefer")
	println("start of check defer")
}
