package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Masukkan kalimat (minimal 3 kata):")

	// buat reader untuk membaca input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// pastiin gaada leading dan trailing whitespace
	input = strings.TrimSpace(input)

	// memecah kalimat menjadi kata-kata
	words := strings.Fields(input)

	// periksa jumlah kata
	if len(words) < 3 {
		fmt.Println("Error: Jumlah kata minimal 3")
		return
	}

	// membalik setiap kata
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}

	// gabungkan kata-kata yang sudah dibalik
	result := strings.Join(words, " ")

	fmt.Println("Hasil pembalikan kata:")
	fmt.Println(result)
}
