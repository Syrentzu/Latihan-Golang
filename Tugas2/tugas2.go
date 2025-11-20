package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	// var kataPertama string
	// var kataKedua string
	// var kataKetiga string
	// var kataKeempat string
	// var kataKelima string

	// kataPertama = "Bootcamp"
	// kataKedua = "Digital"
	// kataKetiga = "Skill"
	// kataKeempat = "Sanbercode"
	// kataKelima = "Golang"

	// fmt.Println(kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat + " " + kataKelima)

	//soal 2
	// halo := "Halo Dunia"
	// var ganti string = "Dunia"
	// var tukar string = "Golang"

	// var hasil string = strings.Replace(halo, ganti, tukar, -1)
	// fmt.Println(hasil)

	//soal 3
	// var kataPertama string = "saya"
	// var kataKedua string = "senang"
	// var kataKetiga string = "belajar"
	// var kataKeempat string = "golang"
	// var ubahKedua string = strings.Replace(kataKedua, "s", "S", 1)
	// var ubahKetiga string = strings.Replace(kataKetiga, "r", "R", 1)
	// var ubahKeempat string = strings.ToUpper(kataKeempat)
	// var hasil string = kataPertama + " " + ubahKedua + " " + ubahKetiga + " " + ubahKeempat
	// fmt.Println(hasil)

	//soal 4
	// var angkaPertama = "8"
	// var angkaKedua = "5"
	// var angkaKetiga = "6"
	// var angkaKeempat = "7"
	// var pertama, _ = strconv.Atoi(angkaPertama)
	// var kedua, _ = strconv.Atoi(angkaKedua)
	// var ketiga, _ = strconv.Atoi(angkaKetiga)
	// var keempat, _ = strconv.Atoi(angkaKeempat)
	// var hasil = pertama + kedua + ketiga + keempat
	// fmt.Println(hasil)

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	var ganti = strings.Replace(kalimat, "halo", "Hi", -1)
	var konver = strconv.Itoa(angka)
	var hasil = `"` + ganti + `"` + " - " + konver
	fmt.Println(hasil)
}
