package main

import (
	"fmt"

	"github.com/nasoundead/LeetCode-Go/solutions"
)

func main() {
	res := solutions.TwoSum([]int{2, 7, 11, 15}, 9)

	fmt.Println(res) // Output: [0, 1]
}
