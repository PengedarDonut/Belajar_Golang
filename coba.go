package main

import "fmt"

func main() {
    // TIPE 1: Deklarasi Tertib (Pake 'var')
    // Bagus buat variabel global atau yang butuh tipe data spesifik
    var nama string = "Pengedar Donut"
    var umur int = 20

    // TIPE 2: Deklarasi Santuy (Short Declaration)
    // Pake ':=', Go otomatis nebak tipe datanya.
    // Cuma bisa dipake di DALAM function.
    hobi := "Push Rank Github"
    isJago := true

    // Printf = Print Format (buat nyelipin variabel ke teks)
    // %s = string, %d = digit (angka), %t = boolean (true/false)
    fmt.Printf("Nama: %s\n", nama)
    fmt.Printf("Umur: %d tahun\n", umur)
    fmt.Printf("Hobi: %s\n", hobi)
    fmt.Printf("Status Jago: %t\n", isJago)
}

 
