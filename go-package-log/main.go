package main

import (
	"log"
	"os"
)

func main() {
	// Contoh penggunaan fungsi-fungsi log
	// log.Print("Contoh log.Print()")
	// log.Printf("Contoh log.Printf() dengan format: %s", "format log")
	// log.Println("Contoh log.Println()")
	// log.Fatal("Contoh log.Fatal()")
	// log.Panic("Contoh log.Panic()")

	// Mengatur level log ke level debug
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.SetPrefix("[DEBUG] ")
	// log.Print("Ini adalah pesan log dengan level DEBUG.")

	// Mengatur level log ke level fatal
	// log.SetFlags(log.LstdFlags)
	// log.SetPrefix("[FATAL] ")
	// log.Fatal("Ini adalah pesan log dengan level FATAL.")

	// Logging ke file
	// Membuat file log baru
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	// Mengatur output log ke file yang telah dibuat
	log.SetOutput(file)

	// Contoh pesan log
	log.Println("Ini adalah pesan log yang ditulis ke dalam file app.log.")
}
