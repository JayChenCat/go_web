package util

import "strconv"

//bool类型转换为int
func ToInt(flag bool) int {
	if flag {
		return 1
	} else {
		return 0
	}
}

//string类型转换为int
func ToIntValue(text string) int {
	if text == "true" {
		return 1
	} else {
		return 0
	}
}

func Tofloat64(text string) float64 {
	v, _ := strconv.ParseFloat(text, 64)
	return v
}
