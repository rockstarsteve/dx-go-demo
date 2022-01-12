package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a = CustmerInt(100)
	fmt.Println("a:", a)
	fmt.Println("a type:", reflect.TypeOf(a))
	fmt.Println("a type:", reflect.TypeOf(&a))

	fmt.Println("===================================================================")
	var b rune
	b = '国'

	fmt.Println("b:", b)
	fmt.Println("b type:", reflect.TypeOf(b))
	b = 20013
	fmt.Println("b:", string(b))
	fmt.Println("===================================================================")

	var tom Person
	tom.name = "tom"
	tom.age = 23
	tom.bobby = []string{"打羽毛球", "打篮球"}
	fmt.Println(tom)

	fmt.Println("=====================匿名结构体==============================================")
	var s struct {
		name string
		age  int
	}
	s.name = "miller"
	s.age = 20
	fmt.Println("s:", s)

	fmt.Println("==================================================================")

	fp(tom)
	fmt.Println("tom:", tom)
	fp2(&tom)
	fmt.Println("tom:", tom)

	fmt.Println("===================new 是创建指针对象  make是给特点类型初始化返回对象===============================================")

	var per = new(Person)
	per.age = 12
	per.name = "jack"	

	fmt.Println("per:", per)

}

type Person struct {
	name  string
	age   int
	bobby []string
}

type CustmerInt int

func fp(x Person) {
	x.name = "修改的人名"
}

func fp2(x *Person) {
	x.name = "修改的人名"
}
