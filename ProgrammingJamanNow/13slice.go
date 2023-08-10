package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	} // tanda [...] menggantikan jumlah/panjang ARRAY bila kita tidak tahu akan ada berapa panjang array yang akan di buat

	// slice/ambil data dr low hingga sebelum high [low:high]
	// slice/ambil data dr low hingga data terakhir [low:]
	// slice/ambil data dr index 0 / awal hingga sebelum high [:high]
	// slice/ambil data dr index 0 / awal hingga terakhir [:]
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // mencari panjang
	fmt.Println(cap(slice1)) // mencari kapasitas

	// ==================================
	// HATI-HATI SAAT MENGGANTI ISI ARRAY ATAU SLICE KARENA AKAN MENGGANTI ARRAY AWAL

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei diubah"
	// fmt.Println(months)
	// ==================================
	// var slice2 = months[2:4] // hasil : [Maret April] dengan kapasitas 10 (masih ada 8 yg kosong)
	// fmt.Println(slice2)
	// fmt.Println(cap(slice2))

	// var slice3 = append(slice2,"'Fandi'") //apped = menambahkan data ke slice2 (bukan array-nya) tanpa menambah kapasitas
	// fmt.Println(slice3)
	// slice3[1] = "'Bukan Desember'"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(months)
	// ==================================
	//TAPI JIKA KITA APPED (TAMBAH DATA BARU) NAMUN KAPASITAS TIDAK CUKUP MAKA TERBENTUK ARRAY BARU TANPA MENGUBAH ARRAY ASAL
	var slice2 = months[10:] // hasil : [Maret April] dengan kapasitas 10 (masih ada 8 yg kosong)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	var slice3 = append(slice2, "'Fandi'") //apped = menambahkan data ke slice2 (bukan array-nya) tanpa menambah kapasitas
	fmt.Println(slice3)
	slice3[1] = "'Bukan Desember'"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// membuat slice baru dengan MAKE SLICE (rekomendasi diterapkan)
	newSlice := make([]string, 2, 5) // tanda [] menggantikan jumlah/panjang SLICE bila kita tidak tahu akan ada berapa panjang SLICE yang akan di buat

	newSlice[0] = "Fandi"
	newSlice[1] = "Putro"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice harus dengan ukuran yg sama
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)
}
