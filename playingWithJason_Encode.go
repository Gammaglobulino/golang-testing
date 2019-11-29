package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
	Addrs []string
}

func main() {
	p1 := person{
		First: "Andrea",
		Last:  "Mazzanti",
		Age:   49,
		Addrs: []string{
			"Via Panoramica 57 Ponte a Poppi",
			"Via al Forte 2 Lugano",
		},
	}
	p2 := person{
		First: "Stefano",
		Last:  "Mazzanti",
		Age:   53,
		Addrs: []string{
			"Via Panoramica 57 Ponte a Poppi",
			"Via xxx Treviso",
		},
	}
	persons := []person{p1, p2}
	err := json.NewEncoder(os.Stdout).Encode(persons)
	if err != nil {
		fmt.Println(err)
	}
}
