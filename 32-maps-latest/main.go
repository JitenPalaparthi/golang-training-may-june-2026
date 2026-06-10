package main

import "fmt"

func main() {
	map1 := make(map[string]any)
	map1["695011"] = "Trivandrum1" // assign a value to the map
	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560054"] = "Bangalore-3"
	map1["560035"] = "Bangalore-4"
	map1["522001"] = "Guntur-1"
	map1["522002"] = "Guntur-2"
	// generates hash from the key --> 695011 , 64bit hash (8 byte hash)
	println("Len of map1:", len(map1))

	clear(map1)

	fmt.Println(map1)

}

// 1.24+> it uses swiss table
// maps header
// count = 0
// seed = random
// table = nil

// each table contains Groups

// map1["695011"] = "Trivandrum1" // assign a value to the map

// generates hash from the key --> 695011 , 64bit hash (8 byte hash)
// Map
// Table - 0
// Group - 0
// 8 slots

// Hash --> 0x1100aebc101ff fe (64bit)
// H1(57) and H2(7)
// it finds the table number based on h1
// table index = h1 % 4 (assume there are 4 tables), % is evaluated by bitwise operation
// 37%4 = 1 (Table index)
// groupIndex = H1 % 8 (8 is the groups count)
// assume H1 is 37 % 8 = 5
// ControlByte is an array -- 1 byte (8 bits)
// KV is a parallel array
// K -> String
// V -> Any

/* struct Group{
ControlByte uint64 or [8]byte
[3F E E E E E E E] // SIMD 8 bytes are sent at once processor and get resule
Keys [8] string
Values[8] any
}*/
