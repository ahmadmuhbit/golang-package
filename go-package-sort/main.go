package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person // Alias

func (p ByAge) Len() int {
	return len(p)
}

func (p ByAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p ByAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	persons := []Person{
		{"Jokoy", 35},
		{"Ubit", 30},
		{"Bambang", 40},
		{"Citra", 28},
	}

	// Mengurutkan slice elemen numerik dari tipe data int
	numbers := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("Sebelum diurutkan:", numbers)

	sort.Ints(numbers)
	fmt.Println("Setelah diurutkan:", numbers)

	// Mengurutkan slice elemen string
	names := []string{"Zaka", "Ahmad", "Budi", "Eva", "Charles"}
	fmt.Println("Sebelum diurutkan:", names)

	sort.Strings(names)
	fmt.Println("Setelah diurutkan:", names)

	// Memeriksa apakah slice sudah terurut atau belum ?
	fmt.Println("Apakah slice sudah terurut?", sort.IsSorted(sort.IntSlice(numbers)))
	fmt.Println("Apakah slice sudah terurut?", sort.IsSorted(sort.StringSlice(names)))

	// Mengurutkan slice dengan elemen tipe data struct
	fmt.Println("Sebelum diurutkan:", persons)

	sort.Sort(ByAge(persons))
	fmt.Println("Setelah diurutkan:", persons)

	// Mengurutkan slice dengan elemen tipe data struct descending
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age > persons[j].Age
	})
	fmt.Println("Setelah diurutkan secara descending berdasarkan usia:", persons)
}
