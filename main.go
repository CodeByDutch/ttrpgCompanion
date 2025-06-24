package main

import (
	"fmt"
	"ttrpgCompanion/dice"
)

// Code here is strictly for testing purposes atm.

func main() {
	newDice := dice.Dice{Amount: 4, Sides: 6}

	newDice.Roll()
	newDice.Roll()
	newDice.Roll()
	newDice.Roll()
	newDice.Roll()
	newDice.Roll()
	newDice.Roll()

	newDice.DisplayRolls()
	fmt.Println()
	newDice.DisplayRolls(2)
	fmt.Println()
	newDice.DisplayRolls(4)
	fmt.Println()
	newDice.DisplayRolls(-1)

}
