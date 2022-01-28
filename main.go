package main

import (
	"fmt"
	"rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())

	state.PutInBoat()
	fmt.Println(state.ViewState())

	state.CrossRiver()
	fmt.Println(state.ViewState())
}
