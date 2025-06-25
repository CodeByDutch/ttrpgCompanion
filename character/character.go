package character

import (
	"fmt"
	"slices"
	"strings"
)

type Character struct {
	Name       string
	Class      string
	Level      int
	Background string
	Alignment  string
	Experience int
	PlayerName string
	Skills     map[string]int
}

var availableClasses = []string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"}

// ***TO BE IMPLEMENTED AT A LATER TIME***
// var availableSubClasses = map[string][]string{
// 	"Barbarian": {"Path of the berserker", "Path of the totem Warrior"},
// 	"Bard":      {"College of lore", "College of valor"},
// 	"Cleric":    {"Knowledge domain", "Life domain", "Light domain", "Nature domain", "Tempest domain", "Trickery domain", "War domain"},
// 	"Druid":     {"Circle of the land", "Circle of the moon"},
// 	"Fighter":   {"Champion", "Battle master", "Eldritch knight"},
// 	"Monk":      {"Way of the open hand", "Way of the shadow", "Way of the four elements"},
// 	"Paladin":   {"Oath of devotion", "Oath of the ancients", "Oath of vengence"},
// 	"Ranger":    {"Hunter, Beast master"},
// 	"Rogue":     {"Thief", "Assassin", "Arcane trickster"},
// 	"Sorcerer":  {"Draconic bloodline", "Wild magic"},
// 	"Warlock":   {"The archfey", "The fiend", "The great one"},
// 	"Wizard":    {"School of abjuration", "School of conjuration", "School of divination", "School of enchantment", "School of evocation", "School of illusion", "School of necromancy", "School of transmutation"},
// }

var availableRaces = []string{"Dwarf", "Elf", "Halfling", "Human", "Dragonborn", "Gnome", "Half-Elf", "Half-Orc", "Tiefling"}

// ***TO BE IMPLEMENTED AT A LATER TIME***
// var availableSubRaces = map[string][]string{
// 	"Dwarf":      {"Hill", "Mountain"},
// 	"Elf":        {"High", "Wood", "Drow"},
// 	"Halfling":   nil,
// 	"Human":      nil,
// 	"Dragonborn": nil,
// 	"Gnome":      {"Forest", "Rock"},
// 	"Half-Elf:":  nil,
// 	"Half-Orc":   nil,
// 	"Tiefling":   nil,
// }

func (c Character) String() string {
	return fmt.Sprintf("Name: %s\nClass: %s\nExperience: %d\nLevel: %d\n", c.Name, c.Class, c.Experience, c.Level)
}

func CreateCharacter(name, class, race, alignment string) (*Character, error) {
	if !slices.Contains(availableClasses, class) {
		return nil, fmt.Errorf("not an available class, please pick one from the following list:\n[%v]", strings.Join(availableClasses, ", "))
	}

	if !slices.Contains(availableRaces, race) {
		return nil, fmt.Errorf("not an available race, please pick one from the following list:\n[%v]", strings.Join(availableRaces, ", "))
	}

	c := &Character{
		Name:       name,
		Class:      class,
		Alignment:  alignment,
		Level:      1,
		Background: "",
		Experience: 0,
		Skills: map[string]int{
			"Strength":     10,
			"Dexterity":    10,
			"Constitution": 10,
			"Wisdom":       10,
			"Intelligence": 10,
			"Charisma":     10,
		},
	}
	return c, nil
}
