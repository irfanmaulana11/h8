package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// type employe struct {
// 	Name string
// }

type employeNew struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// contoh :  go run main.go 11

func main() {
	arg := os.Args[1]
	argInt, _ := strconv.Atoi(arg)

	employe := []employeNew{
		{Name: "irfan", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Giva", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Yusuf", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Aulia", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Fahmi", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Edi", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Cecep", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Teguh", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Bayu", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
		{Name: "Amali", Alamat: "Jakarta", Pekerjaan: "Developer", Alasan: "Gratis"},
	}

	PrintName(employe, argInt)

}

// PrintName ...
func PrintName(emp []employeNew, index int) {

	if index > len(emp) || index == 0 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	data, _ := json.MarshalIndent(emp[index-1], "", "\t")
	fmt.Print(string(data))

}
