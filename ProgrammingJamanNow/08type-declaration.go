package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpFandi noKTP = "1234567890"
	var marriedStatus Married = true

	fmt.Println(noKtpFandi)
	fmt.Println(marriedStatus)
}
