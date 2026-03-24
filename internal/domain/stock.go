package domain

type Stock struct {
	Symbol string
	Quantity float64
	CurrentPrice float64
}

func (s Stock) Value() float64 {
	return s.Quantity * s.CurrentPrice
}

