package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // artinya i = i + 10
	fmt.Println(i)

	i++ // artinya i + 1
	fmt.Println(i)

	var negatif = -100
	var positif = +100
	fmt.Println(negatif)
	fmt.Println(positif)
}
