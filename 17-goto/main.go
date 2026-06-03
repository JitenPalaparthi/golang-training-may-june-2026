package main

func main() {

	num := 1

loop:
	if num > 10 {
		goto exit
	}
	if num%2 == 0 {
		goto even
	} else {
		goto odd
	}

even:
	println(num, "is even number")
	num++
	goto loop

odd:
	println(num, "is odd number")
	num++
	goto loop

exit:
	println("ending goto ")

}
