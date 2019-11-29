package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}
type secretAgent struct {
	person
	ltk  bool
	guns []string
}

func main() {
	ps1 := person{name: "Andrea", lastname: "Mazzanti", age: 49}
	sa := secretAgent{
		person: ps1,
		ltk:    true,
		guns: []string{
			"rifle",
			"chainsaw",
			"pistol",
		},
	}
	list := map[string]secretAgent{ps1.lastname: sa}
	fmt.Println(list["Mazzanti"])
}
