package math

import "testing"

func Test_Float(t *testing.T) {
	x, y := 1.1234567890123456, 2.2345678901234567
	t.Log("add")
	t.Log(Add(x, y))
	t.Log(Add2(x, y))

	t.Log("sub")
	t.Log(Sub(x, y))
	t.Log(Sub2(x, y))

	t.Log("mul")
	t.Log(Mul(x, y))
	t.Log(Mul2(x, y))

	t.Log("div")
	t.Log(Div(x, y))
	t.Log(Div2(x, y))
}

func Benchmark_Add(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Add(x, y)
	}
}

func Benchmark_Add2(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Add2(x, y)
	}
}

func Benchmark_Sub(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Sub(x, y)
	}
}

func Benchmark_Sub2(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Sub2(x, y)
	}
}

func Benchmark_Mul(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Mul(x, y)
	}
}

func Benchmark_Mul2(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Mul2(x, y)
	}
}

func Benchmark_Div(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Div(x, y)
	}
}

func Benchmark_Div2(b *testing.B) {
	x, y := 1.1234567890123456, 2.2345678901234567
	for n := 0; n < b.N; n++ {
		Div2(x, y)
	}
}
