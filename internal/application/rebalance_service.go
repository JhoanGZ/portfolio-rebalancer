package application

import (
	"portfolio-rebalancer/internal/domain"
	"portfolio-rebalancer/internal/types"
)

type RebalanceService struct{}

// Execute orchestration  for rebalance use case
func (s RebalanceService) Execute (p domain.Portfolio) []types.RebalanceAction {
	return p.Rebalance()
}
