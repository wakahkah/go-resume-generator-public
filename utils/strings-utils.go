package utils

import "strings"

func ArrayToStr(ary []string) string {
	comma := ", "

	str := strings.Join(ary[:], comma)

	lastStr := ReplaceLast(str, comma, " and ")

	return lastStr
}

func ReplaceLast(x, y, z string) (x2 string) {
	i := strings.LastIndex(x, y)
	if i == -1 {
		return x
	}
	return x[:i] + z + x[i+len(y):]
}
