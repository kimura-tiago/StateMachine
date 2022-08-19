package stateMachine

import "fmt"

type BalancedState struct {
	payoutContext *PayoutContext
}

func (i *BalancedState) validation(marketplace, counterParty int) error {
	return fmt.Errorf("Marketplace or Counterparty already validated!")
}

func (i *BalancedState) hasbalance(amount float64) error {
	return fmt.Errorf("balance already validated!")
}
func (i *BalancedState) payCounterparty(amount float64) error {
	i.payoutContext.balance = i.payoutContext.balance - amount
	i.payoutContext.setState(i.payoutContext.success)
	fmt.Printf("Payout finished!\n")
	fmt.Printf("New balance %f\n", i.payoutContext.balance)
	return nil
}
