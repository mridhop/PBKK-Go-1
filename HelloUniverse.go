package main

import "fmt"

func main() {
	var number int

	fmt.Println("Masukkan sebuah angka: ")
	fmt.Scan(&number)

	if number == 42 {
		fmt.Println("Hello Universe!")
	} else {
		fmt.Println(number)
	}
}
