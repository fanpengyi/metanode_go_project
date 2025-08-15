package homework1

/*
*
`PlusOne` 函数用于将一个表示大整数的整型数组加 1，并返回结果数组。实现思路如下：

1. 从数组末尾开始遍历，每一位加 1。
2. 如果当前位小于 9，直接加 1并返回结果（无需进位）。
3. 如果当前位等于 9，加 1后变为 0，继续向前进位。
4. 如果所有位都是 9，最后会新建一个长度加 1的新数组，最高位为 1，其余为 0。

这样可以正确处理各种进位情况，返回加 1后的新数组。
*/
func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 最高位有进位
	result := make([]int, n+1)
	result[0] = 1
	return result
}
