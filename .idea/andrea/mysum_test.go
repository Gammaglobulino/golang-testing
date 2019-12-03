package andrea

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		result int
	}

	var tests = []test{
		{
			data:   []int{21, 21},
			result: 42,
		},
		{
			data:   []int{1, 1},
			result: 2,
		},
	}
	for _, v := range tests {
		res := MySum(v.data...)
		if v.result != res {
			t.Error("Expected", v.result, "got", res)
		}
	}
}
