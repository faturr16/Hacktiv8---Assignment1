package main

import (
	"fmt"
)

func main() {

	fmt.Println("Your Full Name :")
	var nama string
	fmt.Scanln(&nama)

	fmt.Println("Your Age :")
	var umur int
	fmt.Scanln(&umur)

	fmt.Println(nama, umur)
}
