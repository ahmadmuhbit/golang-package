package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pencarian substring
	var contains bool = strings.Contains("Hello World", "Hei")
	fmt.Println(contains)
	fmt.Println(strings.Contains("seafood", "foo"))

	// Mengganti substring
	strReplace := "Hello World World World"
	newStr := strings.Replace(strReplace, "World", "Golang", 2)
	fmt.Println(newStr)

	fmt.Println(strings.ReplaceAll("Hello World World World", "World", "Golang"))

	// Mengubah semua huruf menjadi lowercase/uppercase
	str := "Hello World"
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println(lower)
	fmt.Println(upper)

	// Menghapus spasi di awal dan di akhir string
	strTrim := "     Hello World     "
	fmt.Println(strTrim)
	trimmed := strings.Trim(strTrim, " ")
	fmt.Println(trimmed)
}
