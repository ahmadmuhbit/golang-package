package main

import (
	"fmt"
	"math"
)

func main() {
	// Akar kuadrat
	x := 25.0
	sqrtX := math.Sqrt(x)
	fmt.Println("Akar kuadrat dari", x, "adalah", sqrtX)
	fmt.Println(math.Sqrt(25))

	// Pangkat
	a := 2.0
	b := 3.0
	powAB := math.Pow(a, b)
	fmt.Println(a, "pangkat", b, "adalah", powAB)

	// Pembulatan
	num1 := 3.4
	roundedNum1 := math.Round(num1)
	fmt.Println("Hasil pembulatan dari", num1, "adalah", roundedNum1)
	fmt.Println(math.Ceil(num1))

	num2 := 3.7
	roundedNum2 := math.Round(num2)
	fmt.Println("Hasil pembulatan dari", num2, "adalah", roundedNum2)
	fmt.Println(math.Floor(num2))

	// Max dan Min
	i := 10.0
	j := 7.0
	max := math.Max(i, j)
	fmt.Println("Nilai maksimum antara", i, "dan", j, "adalah", max)
	min := math.Min(i, j)
	fmt.Println("Nilai minimum antara", i, "dan", j, "adalah", min)
}
