package main

import "fmt"

func main(){
	// String
	var nama string = "Andar Pratama"
	var kota string = "Tangerang"

	// Numeric Types
	var umur int = 28
	var panjang float64 = 6.4

	// Boolean
	var isLogin = true;
	var isRegistered = false;

	// Array
	var result [4]int = [4]int{8, 9, 8, 7}
	var values []int = []int{2, 3, 7, 3}

	// Object
	type Person struct {
		name string
		age int
		city string
		isMerried bool
	}

	person := Person{
		name: "Andar Pratama",
		age: 28,
		city: "Tangerang",
		isMerried: true,
	}

	fmt.Println(person)
}