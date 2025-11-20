package main

import (
	"fmt"
)

func main() {
	//soal 1
	// var panjangPersegiPanjang string = "8"
	// var lebarPersegiPanjang string = "5"
	// var konverPanjang, _ = strconv.Atoi(panjangPersegiPanjang)
	// var konverLebar, _ = strconv.Atoi(lebarPersegiPanjang)
	// var kelilingPersegi = 2 * (konverPanjang + konverLebar)
	// fmt.Println(kelilingPersegi)
	// var luasPersegi = konverPanjang * konverLebar
	// fmt.Println(luasPersegi)
	// var alasSegitiga string = "6"
	// var tinggiSegitiga string = "7"
	// var konverAlas, _ = strconv.Atoi(alasSegitiga)
	// var konverTinggi, _ = strconv.Atoi(tinggiSegitiga)
	// var setengah float32 = 0.5
	// var luasSegitiga float32 = setengah * float32(konverAlas) * float32(konverTinggi)
	// var stringLuas = strconv.FormatFloat(float64(luasSegitiga), 'f', 2, 64)
	// fmt.Println(stringLuas)

	//soal 2
	// var nilaiJohn = 80
	// var nilaiDoe = 50

	// if nilai >= 80 {
	// 	fmt.Println("Nilai A")
	// } else if nilai >= 70 && nilai < 80 {
	// 	fmt.Println("Nilai B")
	// } else if nilai >= 60 && nilai < 70 {
	// 	fmt.Println("Nilai C")
	// } else if nilai >= 50 && nilai < 60 {
	// 	fmt.Println("Nilai D")
	// } else {
	// 	fmt.Println("Nilai E")
	// }

	//soal 3
	// var tanggal = 26
	// var bulan = 10
	// var tahun = 2004
	// var konverTanggal = strconv.Itoa(tanggal)
	// var konverTahun = strconv.Itoa(tahun)
	// switch bulan {
	// case 1:
	// 	fmt.Println(konverTanggal + " " + "Januari" + " " + konverTahun)
	// case 2:
	// 	fmt.Println(konverTanggal + " " + "Februari" + " " + konverTahun)
	// case 3:
	// 	fmt.Println(konverTanggal + " " + "Maret" + " " + konverTahun)
	// case 4:
	// 	fmt.Println(konverTanggal + " " + "April" + " " + konverTahun)
	// case 5:
	// 	fmt.Println(konverTanggal + " " + "Mei" + " " + konverTahun)
	// case 6:
	// 	fmt.Println(konverTanggal + " " + "Juni" + " " + konverTahun)
	// case 7:
	// 	fmt.Println(konverTanggal + " " + "Juli" + " " + konverTahun)
	// case 8:
	// 	fmt.Println(konverTanggal + " " + "Agustus" + " " + konverTahun)
	// case 9:
	// 	fmt.Println(konverTanggal + " " + "September" + " " + konverTahun)
	// case 10:
	// 	fmt.Println(konverTanggal + " " + "Oktober" + " " + konverTahun)
	// case 11:
	// 	fmt.Println(konverTanggal + " " + "November" + " " + konverTahun)
	// case 12:
	// 	fmt.Println(konverTanggal + " " + " Desember" + " " + konverTahun)
	// }

	//soal 4
	var tahun int32 = 2004
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Anda adalah generasi Baby Boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Anda adalah generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Anda adalah generasi Y(Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Anda adalah generasi Z")
	} else {
		fmt.Println("Anda adalah generasi Alpha")
	}
}
