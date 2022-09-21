package main

import (
	"fmt"
)

func main() {

	var nama string
	fmt.Println("Nama Kamu:")
	fmt.Scanln(&nama)

	var umur int
	fmt.Println("Umur Kamu:")
	fmt.Scanln(&umur)

	fmt.Println("Nama kamu adalah: ", nama)
	fmt.Println("Umur kamu:", umur)

	if umur < 30 {
		fmt.Println("Kamu masih muda", nama, "Semangat terus ya :)")
	} else if umur <= 50 {
		fmt.Println("Kamu sudah dewasa", nama, "Saat nya berprilaku seperti orang dewasa")
	} else {
		fmt.Println("Kamu sudah lebih dewasa", nama, "Kamu harus mencari kesenagan sendiri")
	}
}
