package domain

import (
	"testing"
	"portfolio-rebalancer/internal/types"
)

func TestRebalance(t *testing.T) {

	portfolio := Portfolio{
		Holdings: map[string]Stock{
			"META": {Symbol: "META", Quantity: 10, CurrentPrice: 60},
			"AAPL": {Symbol: "AAPL", Quantity: 5, CurrentPrice: 80},
		},
		TargetAllocation: map[string]float64{
			"META": 0.4,
			"AAPL": 0.6,
		},
	}

	actions := portfolio.Rebalance()

	if len(actions) != 2 {
		t.Fatalf("Se esperaban 2 acciones, se obtuvieron %d", len(actions))
	}

	expected := map[string]types.ActionType{
		"META": types.Sell,
		"AAPL": types.Buy,
	}

	for _, action := range actions {
		if action.Action != expected[action.Symbol] {
			t.Errorf("Acción inesperada para %s", action.Symbol)
		}
	}
}

