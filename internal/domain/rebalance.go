package domain

import "portfolio-rebalancer/internal/types"

func (p Portfolio) Rebalance() []types.RebalanceAction {
	actions := []types.RebalanceAction{}

	total := p.TotalValue()

	for symbol, targetPercent := range p.TargetAllocation {

		currentValue := 0.0
		if stock, ok := p.Holdings[symbol]; ok {
			currentValue = stock.Value()
		}

		targetValue := total * targetPercent
		diff := targetValue - currentValue

		if diff > 0 {
			actions = append(actions, types.RebalanceAction{
				Symbol: symbol,
				Action: types.Buy,
				Amount: diff,
			})
		} else if diff < 0 {
			actions = append(actions, types.RebalanceAction{
				Symbol: symbol,
				Action: types.Sell,
				Amount: -diff,
			})
		}
	}

	return actions
}
