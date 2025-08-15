package homework1

/*
*
`TwoSum` 函数用于在整数数组 `nums` 中找到两个数，使它们的和等于目标值 `target`，并返回这两个数的下标。

核心思路如下：

1. 用哈希表 `m` 存储每个数及其下标，方便快速查找。
2. 遍历数组，每次检查 `target-当前数` 是否已在哈希表中出现过。
3. 如果出现过，说明找到了两个数，返回它们的下标。
4. 如果没有，则把当前数和下标加入哈希表，继续遍历。
5. 如果遍历结束还没找到，返回 `nil`。

这样可以保证时间复杂度为 O(n)。
*/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	//  value,bool = m[key]
	//j, ok := m[target-num]：在哈希表 m 中查找键为 target-num 的值，如果存在，则 ok 为 true，j 是该值的下标。
	//如果找到了，说明当前的 num 和之前的某个数之和等于目标值 target，于是返回这两个数的下标 [j, i]。
	for index, num := range nums {
		if value, ok := m[target-num]; ok {
			return []int{value, index}
		}
		//[value]=index
		m[num] = index
	}
	return nil
}
