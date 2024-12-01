package main

import (
	"fmt"

	"github.com/nasoundead/LeetCode-Go/solutions/0001_TwoSum/twosum"
)

func main() {
	res := twosum.TwoSum([]int{2, 7, 11, 15}, 9)

	fmt.Println(res) // Output: [0, 1]
}
