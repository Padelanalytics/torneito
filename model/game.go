package model

import (
	"strings"
)

type Game struct {
	Country string
	Name    string
	Date    string
}

func (a Game) getTournament() Tournament {
	return Tournament{Country: a.Country, Name: a.Name}
}

func (a Game) Compare(b Game) int {
	r := a.getTournament().Compare(b.getTournament())
	if r != 0 {
		return r
	}
	return strings.Compare(a.Date, b.Date)
}
