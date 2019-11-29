package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string
	LastName string
	Age      int
}

func main() {
	/*p1:=person{
		Name:     "Andrea",
		LastName: "Mazzanti",
		Age:      49,
	}
	p2:=person{
		Name:     "Stefano",
		LastName: "Mazzanti",
		Age:      52,
	}
	people:=[]person{p1,p2}

	bs,err:=json.Marshal(people)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	*/
	s := `[{"Name":"Andrea","LastName":"Mazzanti","Age":49},{"Name":"Stefano","LastName":"Mazzanti","Age":52}]`
	var people []person
	bs := []byte(s)

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range people {
		fmt.Println(v.Name, v.LastName, v.Age)
	}
}
