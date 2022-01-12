package main

import "fmt"

func main() {

	var m1 map[string]int = make(map[string]int)

	fmt.Println(m1 == nil)

	m1["没事"] = 23
	m1["刚刚"] = 93

	fmt.Println("m1:", m1)

	fmt.Println("m1[刚刚]:", m1["刚刚"])
	fmt.Println("m1[刚刚2]:", m1["刚刚2"])

	_, ok := m1["刚刚"]
	if ok {
		fmt.Println("amp中存在对应的key")
	} else {
		fmt.Println("amp中不存在对应的key")
	}

	_, ok2 := m1["刚刚2"]
	if ok2 {
		fmt.Println("amp中存在对应的key")
	} else {
		fmt.Println("amp中不存在对应的key")
	}

	for i, i2 := range m1 {
		fmt.Println("i:", i, "i2", i2)
	}

}
