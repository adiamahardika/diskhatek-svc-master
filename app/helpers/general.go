package helpers

import "strconv"

func StrToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}
