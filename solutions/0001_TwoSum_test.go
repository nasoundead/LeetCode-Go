package solutions

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question{
		{
			para{[]int{3, 2, 4}, 6},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{3, 2, 4}, 5},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans{[]int{1, 3}},
		},

		{
			para{[]int{0, 1}, 1},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 3}, 5},
			ans{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		ans, para := q.ans, q.para
		res := TwoSum(para.nums, para.target)
		if !intSliceEqual(res, ans.one) {
			t.Errorf("【input】:%v       【output】:%v       【expect】:%v\n", para, res, ans.one)
		}
	}
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
