package main

import (
	"container/list"
	"fmt"
)

func main() {
	dll := list.New()

	// Insert elemen ke belakang
	dll.PushBack("A")
	dll.PushBack("B")
	dll.PushBack("C")
	dll.PushBack("D")

	dll.PushFront("Z")

	// Menampilkan elemen-elemen double linked list dari depan ke belakang
	for e := dll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Menampilkan elemen-elemen double linked list dari belakang ke depan
	for e := dll.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// Menampilkan data pertama
	fmt.Println(dll.Front().Value)
	// Menampilkan data terakhir
	fmt.Println(dll.Back().Value)

	// Menampilkan data C dan A dengan Next() dan Prev()
	fmt.Println(dll.Front().Next().Next().Next().Value)
	fmt.Println(dll.Back().Prev().Prev().Prev().Value)

	// Menampilkan data nil
	fmt.Println(dll.Front().Next().Next().Next().Next().Next())
	fmt.Println(dll.Back().Prev().Prev().Prev().Prev().Prev())
	fmt.Println(dll.Front().Prev())
	fmt.Println(dll.Back().Next())
}
