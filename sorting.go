package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	last  string
	age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age > a[j].age }

func main() {

	people := []person{{"Andrea", "Mazzanti", 49}, {"Stefano", "Mazzanti", 52}}

	/*ints:=[]int{2,4,67,3,2,4,67,87}
	strings:=[]string{"Andrea","Luca","Stefano","Beatrice"}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(strings)
	sort.Strings(strings)
	fmt.Println(strings)
	*/
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

}
