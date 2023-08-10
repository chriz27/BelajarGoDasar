package main

import "fmt"

func main() {
	// const firsname string = "Fandi"
	// const lastname = "Putro"
	// const value = 1000

	// multi-constant
	const (
		firsname string = "Fandi"
		lastname        = "Putro"
		value           = 1000
	)

	//firsname = "Cp" tidak dapat diterapkan karena bersifat constant sebelumnya
	fmt.Println(firsname)
	fmt.Println(lastname)
	fmt.Println(value)
}
