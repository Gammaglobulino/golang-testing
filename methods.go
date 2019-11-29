package main

import "fmt"

type person struct {
	name     string
	lastname string
}

type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
	saySomething()
}

func (p *person) speak() {
	fmt.Println(p.name, p.lastname)
}
func (p person) saySomething() {
	fmt.Println(p.name, p.lastname)
}

func foo(h human) {
	switch h.(type) {
	case *person:
		fmt.Println(h.(*person).lastname, h.(*person).name)
	case *secretAgent:
		fmt.Println(h.(*secretAgent).lastname, h.(*secretAgent).name, h.(*secretAgent).ltk)
	}
}

func main() {
	ps := person{
		name:     "Andrea",
		lastname: "Mazzanti",
	}
	sa := secretAgent{
		person: person{
			name:     "Stefano",
			lastname: "Mazzanti",
		},
		ltk: true,
	}

	ps1 := &ps
	ps2 := &sa
	ps2.speak()
	ps1.speak()
	ps.saySomething()
	sa.saySomething()
	//foo(ps)
	//foo(sa)
}
