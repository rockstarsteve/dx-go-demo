package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var p = person{
		"tom",
		23,
	}
	fmt.Println("p:", p)
	fmt.Println("p.name:", p.Name)
	fmt.Println("p.age:", p.Age)

	jsonstr, err := json.Marshal(p)

	if err == nil {
		fmt.Println("err nill")
	} else {
		fmt.Println("err not nill")
	}

	fmt.Println(string(jsonstr))

	jsonstr2 := `{"Name":"tom","age":23}`
	var p2 person

	json.Unmarshal([]byte(jsonstr2), &p2)
	fmt.Println("jsonstr2:", p2)

}

type person struct {
	Name string
	Age  int `json:"age"`
}
