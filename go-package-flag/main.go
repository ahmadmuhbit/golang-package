package main

import (
	"flag"
	"fmt"
)

func main() {
	// Mendefinisikan flag
	var name *string = flag.String("name", "Guest", "Nama pengguna")
	var age *int = flag.Int("age", 0, "Usia pengguna")
	var isAdmin *bool = flag.Bool("admin", false, "Status admin")

	// Parse argumen baris perintah
	flag.Parse()

	// Mengakes nilai flag yang diparsing
	fmt.Println("Name:", *name)
	fmt.Println("Usia:", *age)
	fmt.Println("Admin:", *isAdmin)
}
