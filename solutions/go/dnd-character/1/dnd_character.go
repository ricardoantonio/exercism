package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func rollDie() int {
	return rand.Intn(6) + 1
}

func Ability() int {
	rolls := []int{rollDie(), rollDie(), rollDie(), rollDie()}
	min := rolls[0]
	sum := 0
	for _, roll := range rolls {
		sum += roll
		if roll < min {
			min = roll
		}
	}
	return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	char.Hitpoints = 10 + Modifier(char.Constitution)
	return char
}
