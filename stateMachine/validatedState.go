package stateMachine

import "fmt"

type ValidatedState struct {
	payoutContext *PayoutContext
}

func (i *ValidatedState) validation(marketplace, counterParty int) error {
	return fmt.Errorf("Marketplace or Counterparty already validated!")
}

func (i *ValidatedState) hasbalance(amount float64) error {
	if amount > i.payoutContext.balance {
		i.payoutContext.setState(i.payoutContext.rejected)
		return fmt.Errorf("Balance is not enough!")
	}
	i.payoutContext.setState(i.payoutContext.balanced)
	return nil
}
func (i *ValidatedState) payCounterparty(amount float64) error {
	return fmt.Errorf("Balance must validate first!")
}
