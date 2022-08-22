package main

import (
	"fmt"
)

func main() {
	// name := "irfan"
	// age := 26

	// fmt.Printf("%s %d", name, age)
	// fmt.Println("Hallo World!!")

	PrintNumber()
}

// printNumber ...
func PrintNumber() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Genap")
		} else {
			fmt.Println(i, "Ganjil")
		}
	}

}
