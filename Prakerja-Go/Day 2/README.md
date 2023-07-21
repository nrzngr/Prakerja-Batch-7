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

## Tugas 3
### Buatlah sebuah program untuk menghitung luas bangun trapesium
