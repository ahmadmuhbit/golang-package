package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Membuat direktori
	err := os.Mkdir("mydir", 0755)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Direktori berhasil dibuat.")
	}

	// Membuat file dan menulis ke file
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println("Error,", err)
	} else {
		fmt.Println("Tulisan berhasil ditulis ke file.")
	}

	// Membaca environment variable
	username := os.Getenv("USER")
	fmt.Println("Username:", username)

	// Contoh penggunaan os.Getenv dengan library godotenv
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// now do something with s3 or whatever
	fmt.Println(dbHost)
	fmt.Println(dbPort)
}
