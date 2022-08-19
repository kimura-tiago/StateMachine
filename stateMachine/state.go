package stateMachine

type State interface {
	validation(marketplace, counterParty int) error
	hasbalance(amount float64) error
	payCounterparty(amount float64) error
}
