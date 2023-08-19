package main

import (
	"fmt"
	"time"
)

func main() {
	// Mendapatkan waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini:", now)

	// Mendapatkan komponen-komponen waktu
	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println("Tanggal:", day, month, year)
	fmt.Println("Waktu:", hour, minute, second)

	// Membuat waktu yang ditentukan
	createdDate := time.Date(2020, time.September, 19, 7, 10, 50, 20, time.UTC)
	fmt.Println("Tanggal dibuat:", createdDate)

	// Format waktu ke dalam string
	// layout := time.RFC3339
	// formattedTime := now.Format(layout)
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Waktu yang di format:", formattedTime)

	// Parsing string ke dalam waktu
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", "2023-08-19 07:18:20")
	fmt.Println("Waktu yang di parse:", parsedTime)
}
