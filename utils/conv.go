package utils

import "strconv"

func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func ToFloat64(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return float
}
