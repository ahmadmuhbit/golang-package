package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}

func main() {
	var num int = 42
	var str string = "Hello, Reflect!"

	// Mengambil tipe data dan nilai variabel num
	numType := reflect.TypeOf(num)
	// numValue := reflect.ValueOf(num)

	// Mengambil tipe data dan nilai variabel str
	strType := reflect.TypeOf(str)
	strValue := reflect.ValueOf(str)

	fmt.Println("Tipe data num:", numType)
	// fmt.Println("Nilai num:", numValue)
	fmt.Println("Tipe data str:", strType)
	fmt.Println("Nilai str:", strValue)

	// Memeriksa apakah variabel num merupakan int
	if numType.Kind() == reflect.Int {
		fmt.Println("Num adalah int")
	}

	// Memeriksa apakah variabel str merupakan string
	if strType.Kind() == reflect.String {
		fmt.Println("Str adalah string")
	}

	// Mengubah nilai variabel num menjadi 100
	numValue := reflect.ValueOf(&num).Elem()
	numValue.SetInt(100)
	fmt.Println("Updated value of num:", num)

	p := Person{Name: "Ahmad Muhbit", Age: 30}
	t := reflect.TypeOf(p)

	// Memeriksa setiap field dan tag dari struct Person
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, tag)
	}
}
