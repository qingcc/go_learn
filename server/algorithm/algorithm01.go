package main

import "fmt"

/*
*
todo 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

func main() {
	fmt.Println(bettersingleNumber([]int{2, 2, 1, 3, 1}))
}

func singleNumber(nums []int) int {
	wNum := make(map[int]int)
	for _, num := range nums {
		if _, ok := wNum[num]; ok {
			delete(wNum, num)
		} else {
			wNum[num] = num
		}
	}
	aim := 0
	for _, num := range wNum {
		aim = num
		break
	}
	return aim
}

//todo ^抑或运算符 位运算 原理待确认
func bettersingleNumber(nums []int) int {
	res := 0
	for _, value := range nums {
		res = res ^ value
		fmt.Println("in range:", value, res)
	}
	return res
}

func init() {
	ar := [][]int{}
	fmt.Println(len(ar) == 0)
}
