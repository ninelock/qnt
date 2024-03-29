package math

import "github.com/shopspring/decimal"

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) float64 {
	return a / b
}

func Add2(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Float64()
	return f
}

func Sub2(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Float64()
	return f
}

func Mul2(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Float64()
	return f
}

func Div2(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Float64()
	return f
}
