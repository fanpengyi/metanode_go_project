package homework1

import "fmt"

func IsHuiwen(num int) bool {
	if num < 0 {
		return false
	}

	str := fmt.Sprintf("%d", num)
	left, right := 0, len(str)-1

	for left < right {

		if str[left] != str[right] {
			return false
		}
		left++
		right--

	}

	return true

}
