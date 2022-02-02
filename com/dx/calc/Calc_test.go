package calc

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got == want {
		fmt.Println("ok")
		t.Log("----------???")
	} else {
		t.Error("测试不通过！")
	}
}


func BenchmarkAdd(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Add(4,4)
	}
}