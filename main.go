package main

import (
	statemachine "github/StateMachine/stateMachine"
	"log"
)

func main() {
	//CASE 1 - Perfect Flow
	//borned with created state
	payoutState := statemachine.New(5.5, 1, 100)

	payoutState.IncrementBalance(10)

	// Send event to topic / business logic...

	//change state to validated
	errv := payoutState.Validation(1, 100)
	if errv != nil {
		log.Fatalf(errv.Error())
	}

	// Send event to topic / business logic...

	//change state to balanced
	errhb := payoutState.Hasbalance(5.5)
	if errhb != nil {
		log.Fatalf(errhb.Error())
	}

	// Send event to topic / business logic...

	//change state to sucess
	errp := payoutState.PayCounterparty(5.5)
	if errp != nil {
		log.Fatalf(errp.Error())
	}

	//finished Flow, trying to change to other state
	errr := payoutState.PayCounterparty(55.5)
	if errr != nil {
		log.Fatalf(errr.Error())
	}

	////////////////////////////////////////////////////////////////////////////////

	//CASE 2 - Wrong State Flow
	//borned with created state
	payoutState2 := statemachine.New(5.5, 1, 100)
	payoutState2.IncrementBalance(10)

	// Send event to topic / other action...

	//change state to validated
	errv2 := payoutState2.PayCounterparty(100)
	if errv2 != nil {
		log.Fatalf(errv2.Error())
	}

}
