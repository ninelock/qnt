package entity

type Depth struct {
	Ts   int64
	Cp   string
	Asks [][]float64
	Bids [][]float64
}
