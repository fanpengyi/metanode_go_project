package homework1

func IsValid(str string) bool {

	slice := []rune{}
	//存的尾巴
	mapRune := make(map[rune]rune, 3)
	mapRune[')'] = '('
	mapRune[']'] = '['
	mapRune['}'] = '{'

	for _, ch := range str {
		switch ch {
		case '(', '[', '{':
			slice = append(slice, ch)
		case ')', ']', '}':
			if len(slice) == 0 || slice[len(slice)-1] != mapRune[ch] {
				return false
			}
			slice = slice[:len(slice)-1]
		}

	}
	return len(slice) == 0
}
