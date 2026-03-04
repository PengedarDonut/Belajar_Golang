package main

import "fmt"

func main() {
	// Deklarasi variabel menggunakan var (tipe data) = (nilai)
	var nama string = "Pengedar"
	var nama2 string
	nama2 = "Donut"
	fmt.Printf("Nama Lengkap : %s %s \n", nama, nama2)

	// Deklarasi variabel tanpa tipe data

	nama3 := "Udin"
	nama4 := "Ngaga"
	fmt.Printf("Nama lengkap udin : %s %s \n", nama3, nama4)

	// Deklarasi multi variabel
	var first, second, third string
	first = "Satu"
	second = "Dua"
	third = "Tiga"
	fmt.Printf("Nilai variabel : %s, %s, %s \n", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Printf("Nilai variabel : %s, %s, %s \n", fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Printf("Nilai variabel : %s, %s, %s \n", seventh, eight, ninth)

	satu, hari_ini_kamis, tanggal, bulan := 1, true, 05, "Maret"
	fmt.Printf("Nilai variabel : %d, %t, %d, %s \n", satu, hari_ini_kamis, tanggal, bulan)

	// Deklarasi variabel menggunakan keyword new
	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""
}
