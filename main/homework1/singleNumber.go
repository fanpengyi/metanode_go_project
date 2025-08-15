package homework1

func SingleNumber(nums []int) int {
	countMap := make(map[int]int)

	for _, value := range nums {
		countMap[value]++
	}

	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}
	return -99
}
