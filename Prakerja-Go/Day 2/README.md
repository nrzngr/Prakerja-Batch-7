# Day 2 Go-Lang | Alterra Academy

## Tugas 1
### Buatlah sebuah program untuk menentukan bilang prima
Langkah 1: Membuat Program Dasar
1. Buatlah sebuah folder baru untuk proyek Anda dan masuk ke dalam folder tersebut menggunakan terminal atau command prompt.
2. Buatlah file baru dengan nama main.go di dalam folder proyek Anda.

Langkah 2: Menulis Kode Program
```go
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
```


## Tugas 2
### Buatlah sebuah program untuk menentukan bilangan kelipatan 7
Langkah 1: Menulis Kode Program
```go
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
```


## Tugas 3
### Buatlah sebuah program untuk menghitung luas bangun trapesium
Langkah 1: Menulis Kode Program
Rumus Luas Trapesium: ``` luas = 0.5 * (sisi1 + sisi2) * tinggi ```
```go
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
```
