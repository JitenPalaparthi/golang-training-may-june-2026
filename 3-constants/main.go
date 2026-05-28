package main

const (
	MAX, MIN     = 999, 1
	TICKS    int = 60
	SECOND       = 1
	MINUTE       = SECOND * TICKS // Is it possible ? as long as the right handside are also constatns
	HOUR         = MINUTE * TICKS
	DAY          = HOUR * 24
	A, B         = 10, 20
	SOMEC        = (A+B/2*B)-(A+3)+A%B+343+DAY > 10000 && true // does not make sense of this expression, but

	//SOMEC1 = true + 10

	// What is this evaluated? Not at runtime, at compiletime
)

func main() {
	const PI float32 = 3.14

	//MAX += 1 // This is not possible

	// min := 60

	// const SOMTING = MINUTE * min // This is not possible

	println(MAX, SOMEC)

	PrintHourSec()

	//num := SOMEC * 10

}

// Compitime vs Runtime
