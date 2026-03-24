package domain

type Portfolio struct {
	Holdings map[string]Stock
	TargetAllocation map[string]float64
}


func (p Portfolio) TotalValue() float64{
	total := 0.0
	for _, stock := range p.Holdings {
		total += stock.Value()
	}
	return total
}

