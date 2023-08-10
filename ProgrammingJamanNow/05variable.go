package main

import "fmt"

func main() {
	var name string

	name = "Fandi"
	fmt.Println(name)

	name = "Fandi Chriswantoro"
	fmt.Println(name)

	// penulisan tanpa menyebutkan type data
	var friendname = "Budi"
	fmt.Println(friendname)

	var age = 30
	fmt.Println(age)

	// keyword "var" tidak wajib, namun wajib menggunakan titik dua samadengan ":=" di awal init
	country := "Indonesia" // init var country
	fmt.Println(country)

	country = "Amerika" // tanpa titik dua hanya samadengan krn sudah di init di line 22
	fmt.Println(country)

	// multi-variable
	var (
		firsname = "Fandi"
		lastname = "Putro"
	)
	fmt.Println(firsname)
	fmt.Println(lastname)
}
