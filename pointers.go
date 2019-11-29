package main

import "fmt"

type person struct {
	name     string
	lastname string
}

func main() {
	p := person{
		name:     "Andrea",
		lastname: "Mazzanti",
	}
	changeme(&p, "Stefano", "Mazzanti")
	fmt.Println(p.name, p.lastname)
}

func changeme(p *person, name string, lastname string) {
	(*p).name, (*p).lastname = name, lastname

}
