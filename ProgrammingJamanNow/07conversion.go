package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //nilia tidak mampu ditampung, maka akan melooping ke negatif-positif

	// konversi huruf - byte - string
	var name = "Eko"
	var e byte = name[0]           // ambil karakter ke-1 dari var name masukkan ke var e
	var eString string = string(e) // konversi nilai pada var e ke string dan masukkan ke var eString
	fmt.Println(name)
	fmt.Println(eString)

}
