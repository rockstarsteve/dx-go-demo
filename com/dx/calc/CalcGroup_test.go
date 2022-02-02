package calc

import (
	"testing"
)



func TestCalcGroup(t *testing.T) {
	type testCase struct {
		a,b,want int

	}
	testGroup :=map[string]testCase{
		"test1": {1,2,3},
		"test2": {1,3,4},
		"test3": {1,4,5},
	}
	for i, c := range testGroup {
		t.Run(i, func(t *testing.T) {
			got := Add(c.a,c.b)
			if c.want != got {
				t.Fatalf("测试不通过%v",c)
			}
		})

	}
}
