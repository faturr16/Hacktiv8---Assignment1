package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama, alamat, pekerjaan, alasan string
}

func main() {

	murid := []student{
		{
			nama:      "Nama: Fatur\n",
			alamat:    "Alamat: Depok\n",
			pekerjaan: "Pekerjaan: Mahasiswa\n",
			alasan:    "Alasan memilih golang:ingin menambah ilmu dan mendalami tentang backend\n",
		},

		{
			nama:      "Nama: Alex\n",
			alamat:    "Alamat: Demak\n",
			pekerjaan: "Pekerjaan: Akunta\n",
			alasan:    "Alasan memilih golang: bahasa yang ringan untuk laptop kentang \n",
		},

		{
			nama:      "Nama: RObin\n",
			alamat:    "Alamat: Surabaya\n",
			pekerjaan: "Pekerjaan: sales\n",
			alasan:    "Alasan: Keren aja\n",
		},

		{
			nama:      "Nama: Josh\n",
			alamat:    "Alamat: Bogor\n",
			pekerjaan: "Pekerjaan: Mahasiswa\n",
			alasan:    "Alasan: mau nyaingin bjorka\n",
		},

		{
			nama:      "Nama: George\n",
			alamat:    "Alamat: Bogor\n",
			pekerjaan: "Pekerjaan: Chef\n",
			alasan:    "Alasan memilih golang: bahasa pemrogaraman yang ringan \n",
		},

		{
			nama:      "Nama: Marvin\n",
			alamat:    "Alamat: Cipete\n",
			pekerjaan: "Pekerjaan: Mahasiswa\n",
			alasan:    "Alasan memilih golang: bahasa yang ringan untuk laptop kentang \n",
		},

		{
			nama:      "Nama: Kevin\n",
			alamat:    "Alamat: Alaska\n",
			pekerjaan: "Pekerjaan: Mahasiswa\n",
			alasan:    "Alasan: mau nyaingin bjorka\n",
		},

		{
			nama:      "Nama: Edward\n",
			alamat:    "Alamat: California\n",
			pekerjaan: "Pekerjaan: Sales\n",
			alasan:    "Alasan memilih golang: Ingin mendalamin tentang Golang \n",
		},

		{
			nama:      "Nama: Thomas\n",
			alamat:    "Alamat: Kemang\n",
			pekerjaan: "Pekerjaan:Polisi \n",
			alasan:    "Alasan memilih golang: biar keliatan bisa ngoding \n",
		},

		{
			nama:      "Nama: Aswin\n",
			alamat:    "Alamat: Bogor\n",
			pekerjaan: "Pekerjaan: Chef\n",
			alasan:    "Alasan memilih golang: mau nyaigin bjorka \n",
		},

		{
			nama:      "Nama: Dion\n",
			alamat:    "Alamat: Depok\n",
			pekerjaan: "Pekerjaan: Mahasiswa\n",
			alasan:    "Alasan memilih golang: Keren dan Ringan \n",
		},
	}

	absen := os.Args[1]
	for i, v := range murid {
		indexTostring := strconv.Itoa(i)
		if indexTostring == absen {
			fmt.Println(v)
		}

	}

}
