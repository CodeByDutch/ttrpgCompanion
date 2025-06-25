package main

import (
	// "fmt"
	// "ttrpgCompanion/dice"

	"fmt"
	"ttrpgCompanion/character"
)

// Code here is strictly for testing purposes atm.

func main() {
	// newDice := dice.Dice{Amount: 4, Sides: 6}

	// newDice.Roll()
	// newDice.Roll()
	// newDice.Roll()
	// newDice.Roll()
	// newDice.Roll()
	// newDice.Roll()
	// newDice.Roll()

	// newDice.DisplayRolls()
	// fmt.Println()
	// newDice.DisplayRolls(2)
	// fmt.Println()
	// newDice.DisplayRolls(4)
	// fmt.Println()
	// newDice.DisplayRolls(-1)

	newChar, err := character.CreateCharacter("Dutch", "Ranger", "Demi-God", "Divine")
	if err != nil {
		fmt.Printf("Error creating character: %v\n", err)
		return
	}

	fmt.Println(newChar)
}
