package main

func main() {

	defer func() { // func1
		defer println("End of main")
		func() { //func1.1
			defer func() { // func1.1.1
				defer println("Again end of main")
			}()
		}()
		// all defers inside the func are called
	}() ///

	// defer func() {
	// 	println("Again end of main")
	// }()

	defer println("Start of main")

	// time to pop deferred calls and execute

}

// defer defers the execution of a function/method to the end of the call stack
// all defers for caller are stored in stack , and then when it is the end of the call stack they are poped and executed
