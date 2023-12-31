package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah sebuah bilangan merupakan bilangan prima
func bilPrim(){
	var primNum, primcount int
	
	primcount = 0

	fmt.Print("Masukan angka = ")
	fmt.Scanln(&primNum)

	for i := 2; i < primNum/2; i++ {
		if primNum % i == 0 {
			primcount++
			break
		}
	}

	if primcount == 0 && primNum != 1 {
		fmt.Println(primNum, " adalah bilangan prima")
	} else {
		fmt.Println(primNum, " bukan bilangan prima")
	}
}

// Fungsi untuk mengecek suatu bilangan habis dibagi 7 atau tidak
func kel7(){
	var num int

	// Untuk menerima user input dan meletakkan hasil input ke variabel num
    fmt.Print("Masukan angka =  ")
    fmt.Scanln(&num)

	// If statement jika bilangan tersebut habis dibagi 7 atau tidak
    if num % 7 == 0 {
        fmt.Println(num, " Merupakan bilangan kelipatan 7")
    } else {
        fmt.Println(num, " Bukan merupakan bilangan kelipatan 7")
    }
}

// Fungsi untuk menghitung luas trapesium
func trapesium(){

	var sisi1, sisi2, tinggi, luas float64
	

	fmt.Print("Masukan sisi sejajar pertama = ")
	fmt.Scanln(&sisi1)

	fmt.Print("Masukan sisi sejajar kedua = ")
	fmt.Scanln(&sisi2)

	fmt.Print("Masukan tinggi trapesium = ")
	fmt.Scanln(&tinggi)

	luas = 0.5 * (sisi1 + sisi2) * tinggi
	fmt.Print("Luas trapesium nya adalah = ", luas)
}

func main(){
	fmt.Println("-----------Prime---------------")
	bilPrim()
	fmt.Println()

	fmt.Println("----------DivisibleBy7-----------")
	kel7()
	fmt.Println()

	fmt.Println("----------Trapesium--------------")
	trapesium()
	fmt.Println()
}
