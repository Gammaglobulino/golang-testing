package main

import "fmt"

func main() {
	x := []int{1, 3, 4, 5, 6, 7, 8, 9, 10} //slice made by composite literals
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	mp := map[string]string{
		"Andrea":  "Mazzanti",
		"Stefano": "Mazzanti",
		"Luca":    "Mazzanti"}
	fmt.Println(mp)
	i := 0
	for k, v := range mp {
		fmt.Println(i, k, v)
		i++
	}

}
