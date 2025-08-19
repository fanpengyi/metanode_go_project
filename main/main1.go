package main

import (
	"base_code/main/homework1"
	"fmt"
)

func main1() {
	// 1 只出现一次的数字
	nums := []int{2, 2, 1}
	fmt.Println(homework1.SingleNumber(nums)) // 输出: 1
	// 2 回文数
	num := 12131
	fmt.Println(homework1.IsHuiwen(num)) // 输出: true
	// 3 有效的括号
	fmt.Println(homework1.IsValid("()"))     // 输出: true
	fmt.Println(homework1.IsValid("()[]{}")) // 输出: true
	fmt.Println(homework1.IsValid("(]"))     // 输出: false
	fmt.Println(homework1.IsValid("([)]"))   // 输出: false
	fmt.Println(homework1.IsValid("{[]}"))   // 输出: true
	//4 最长公共前缀
	fmt.Println(homework1.LongestCommonPrefix([]string{"flower", "flow", "flight"})) // 输出: "fl"
	fmt.Println(homework1.LongestCommonPrefix([]string{"dog", "racecar", "car"}))    // 输出: ""
	//5 两数之和
	fmt.Println(homework1.PlusOne([]int{1, 2, 3})) // 输出: [1 2 4]
	fmt.Println(homework1.PlusOne([]int{9, 9, 9})) // 输出: [1 0 0 0]
	//6 删除排序数组中的重复项
	nums6 := []int{1, 1, 2, 2, 3}
	k := homework1.RemoveDuplicates(nums6)
	fmt.Println(k)         // 输出: 3
	fmt.Println(nums6[:k]) // 输出: [1 2 3]
	// 7 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := homework1.Merge(intervals)
	fmt.Println(merged) // 输出: [[1 6] [8 10] [15 18]]
	// 8 两数之和
	nums2 := []int{2, 7, 11, 15}
	target := 9
	result := homework1.TwoSum(nums2, target)
	fmt.Println(result) // 输出: [0 1]

}
