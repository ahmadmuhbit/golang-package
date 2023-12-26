package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Membuat regexp object
	regex := regexp.MustCompile("[0-9]+")

	// Mencari kecocokan regexp
	matched := regex.MatchString("Halo 123 dunia")
	fmt.Println("Apakah terdapat angka dalam string?", matched)

	// Mencari semua angka dalam string
	matches := regex.FindAllString("123abc456", -1)
	fmt.Println("Angka-angka dalam string:", matches)

	// Mengganti teks dengan regexp
	newString := regex.ReplaceAllString("Halo 123 dunia", "angka")
	fmt.Println("String setelah diubah:", newString)

	// Memvalidasi teks dengan regexp
	isValid := regex.MatchString("12345")
	fmt.Println("Apakah 12345 kode pos yang valid", isValid)

	// Contoh program validasi email
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := "example@email.com"

	match, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Email is valid:", match)
}
