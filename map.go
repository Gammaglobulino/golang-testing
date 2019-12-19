package main

import "fmt"

func main() {
	//map definition
	m := map[string]int{
		"Andrea":  1,
		"Stefano": 2,
		"Luca":    3,
	}

	fmt.Println(m)

	if v, ok := m["Andrea"]; ok {
		fmt.Printf("%d is there", v)
	}
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)
	}

}
