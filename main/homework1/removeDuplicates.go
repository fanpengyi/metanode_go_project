package homework1

/*
*
`RemoveDuplicates` 函数用于原地删除非严格递增数组中的重复元素，使每个元素只出现一次，并返回唯一元素的个数。实现思路如下：

1. 用变量 `k` 记录下一个唯一元素要放的位置，初始为 1。
2. 遍历数组，如果当前元素与前一个唯一元素不同，则将其放到 `k` 位置，并 `k++`。
3. 遍历结束后，`k` 就是唯一元素的数量，且数组前 `k` 个位置存放着所有唯一元素。

这样可以保证空间复杂度为 O(1)，且元素顺序不变。
}
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
