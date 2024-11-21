package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Selamat datang di Kalkulator Sederhana!")
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	result := calculate(num1, num2, operator)
	if result != nil {
		fmt.Printf("Hasil: %.2f\n", *result)
	} else {
		fmt.Println("Operator tidak valid.")
	}
}

func calculate(a, b float64, operator string) *float64 {
	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			return nil
		}
	default:
		return nil
	}
	return &result
}
