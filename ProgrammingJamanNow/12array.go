package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Fandi"
	names[1] = "Chriswantoro"
	names[2] = "Putro"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// atau dengan format
	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//mencari panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]int // dapat mengecek panjang array walau masih kosong
	fmt.Println(len(lagi))
}
