package main

import "fmt"

func main() {

	//  接口存动态类型和动态值

	var p personsss
	var d dog


	animal(p)
	animal(d)
}

func animal(r runner) {
	r.run()
}

func (p personsss) run() {
	fmt.Println("人在跑。。。")
}
func (a dog) run() {
	fmt.Println("动物在跑。。。")
}

type runner interface {
	run()
}

type personsss struct {
}
type dog struct {
}
