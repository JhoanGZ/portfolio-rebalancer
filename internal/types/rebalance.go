package types

type ActionType int

const (
	 Hold ActionType = iota 
	 Buy
	 Sell
)

type RebalanceAction struct {
	Symbol string
	Action ActionType
	Amount float64
}



