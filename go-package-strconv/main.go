package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi string ke integer
	num1, err := strconv.Atoi("95")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(num1)
		fmt.Printf("Tipe data num: %T\n", num1)
	}

	// Konversi integer ke string
	str1 := strconv.Itoa(95)
	fmt.Println(str1)
	fmt.Printf("Tipe data str1: %T\n", str1)

	// Konversi string ke floating point
	num2, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(num2)
		fmt.Printf("Tipe data num2: %T\n", num2)
	}

	// Konversi floating point ke string
	str2 := strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str2)
	fmt.Printf("Tipe data str2: %T\n", str2)

	// Konversi string ke boolean
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(boolean)
		fmt.Printf("Tipe data boolean: %T\n", boolean)
	}

	// Konversi boolean ke string
	str3 := strconv.FormatBool(false)
	fmt.Println(str3)
	fmt.Printf("Tipe data str3: %T\n", str3)
}
