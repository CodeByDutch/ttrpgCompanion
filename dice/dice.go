package dice

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Amount int
	Sides  int
	Rolls  [][]int
	Sums   []int
}

func (d *Dice) Roll() []int {
	roll := make([]int, 0, d.Amount)
	sum := 0
	for i := 0; i < d.Amount; i++ {
		val := rand.Intn(d.Sides) + 1
		roll = append(roll, val)
		sum += val
	}
	d.Rolls = append(d.Rolls, roll)
	d.Sums = append(d.Sums, sum)
	return roll
}

func (d Dice) DisplayRolls(amount ...int) {
	total := len(d.Rolls)

	n := total

	if len(amount) > 0 { // Checks to see if the user gave atleast one value as a parameter
		if amount[0] < 0 { // If the first value given is < 0, the program returns
			return
		} else if amount[0] > 0 && amount[0] < total { // If the first parameter given is > 0, and < the total amount of Rolls
			n = amount[0] // the parameter given is then set to the n variable
		}
	}

	start := total - n

	for i := start; i < total; i++ {
		fmt.Println("Roll ->", d.Rolls[i], "\tSum ->", d.Sums[i])
	}
}
