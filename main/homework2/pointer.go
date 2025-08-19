package homework2

func Add(num *int) {
	*num += 10
}

func ChengArray(arr *[]int) {
	for index, _ := range *arr {
		(*arr)[index] *= 2
	}

}
