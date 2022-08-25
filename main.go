package main

import (
	"fmt"
)

type employe struct {
	Name string
}

func main() {
	//PrintName()
	es := []*employe{
		{Name: "irfan"},
		{Name: "Giva"},
		{Name: "Yusuf"},
		{Name: "Aulia"},
		{Name: "Fahmi"},
		{Name: "Edi"},
		{Name: "Cecep"},
		{Name: "Teguh"},
		{Name: "Bayu"},
		{Name: "Amali"},
	}

	var _ = func(em []*employe) *string {
		for _, e := range em {
			fmt.Println(e.Name)
		}
		return nil
	}(es)

}

// PrintName ...
func PrintName() {
	var names1 = []string{"irfan", "Aulia", "San", "Wicak", "Iqbal"}
	var names2 = []string{"Giva", "Fahmi", "Rizki", "Billy", "Fajar"}
	names1 = append(names1, names2...)

	for _, n := range names1 {
		fmt.Println(n)
	}
}
