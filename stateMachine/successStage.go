package stateMachine

import "fmt"

type SuccessState struct {
	payoutContext *PayoutContext
}

func (i *SuccessState) validation(marketplace, counterParty int) error {
	return fmt.Errorf("Marketplace or Counterparty already validated!")
}

func (i *SuccessState) hasbalance(amount float64) error {
	return fmt.Errorf("balance already validated!")
}
func (i *SuccessState) payCounterparty(amount float64) error {
	return fmt.Errorf("balance already paid!")
}
