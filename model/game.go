package model

import (
	"strings"
)

// Game represents tournament game and a csv line in the file
type Game struct {
	Country  string
	Name     string
	Serie    string
	Division string
	Date     string
	Round    string
	Category string
	Teams    uint8
	Players  []string
	Sets     uint8
	Scores   []int8
}

// Compare returns -1 if g is less than other, 0 if equals and +1 if greater than other
func (g Game) Compare(other Game) int {
	// equals
	if g.Country == other.Country &&
		g.Name == other.Name &&
		g.Serie == other.Serie &&
		g.Division == other.Division &&
		g.Category == other.Category &&
		g.Round == other.Round &&
		g.Date == other.Date {
		return 0
	}
	// less than
	if g.Country < other.Country ||
		g.Country == other.Country && g.Name < other.Name ||
		g.Country == other.Country && g.Name == other.Name && g.Serie < other.Serie ||
		g.Country == other.Country && g.Name == other.Name && g.Serie == other.Serie && g.Division < other.Division ||
		g.Country == other.Country && g.Name == other.Name && g.Serie == other.Serie && g.Division == other.Division && cCategory[g.Category] < cCategory[other.Category] ||
		g.Country == other.Country && g.Name == other.Name && g.Serie == other.Serie && g.Division == other.Division && cCategory[g.Category] == cCategory[other.Category] && cRound[g.Round] < cRound[other.Round] ||
		g.Country == other.Country && g.Name == other.Name && g.Serie == other.Serie && g.Division == other.Division && cCategory[g.Category] == cCategory[other.Category] && cRound[g.Round] == cRound[other.Round] && strings.Compare(g.Date, other.Date) < 0 {
		return -1
	}
	return 1
}

// IsTournament returns true if the game is part of the tournament
func (a Game) IsTournament(t Tournament) bool {
	return a.Country == t.Country &&
		a.Name == t.Name &&
		a.Serie == t.Serie &&
		a.Division == t.Division
}
