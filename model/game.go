package model

import (
	"strconv"
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
// The order is: Country, Name, Serie, Division, Category, Round, Date
// Players and Scores are not compared
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
func (g Game) IsTournament(t Tournament) bool {
	return g.Country == t.Country &&
		g.Name == t.Name &&
		g.Serie == t.Serie &&
		g.Division == t.Division
}

// NewFromRecord creates a new Game from a csv line in this case a []string.
// It is the opposite of ToRecord.
func NewFromRecord(r []string) Game {
	sets, _ := strconv.Atoi(r[cR["sets"]])
	teams, _ := strconv.Atoi(r[cR["teams"]])
	return Game{
		Country:  r[cR["country"]],
		Name:     r[cR["name"]],
		Serie:    r[cR["serie"]],
		Division: r[cR["division"]],
		Date:     r[cR["date"]],
		Round:    r[cR["round"]],
		Category: r[cR["category"]],
		Teams:    uint8(teams),
		Players:  []string{r[cR["p1_last"]], r[cR["p1_first"]], r[cR["p2_last"]], r[cR["p2_first"]], r[cR["p3_last"]], r[cR["p3_first"]], r[cR["p4_last"]], r[cR["p4_first"]]},
		Sets:     uint8(sets),
		Scores:   ScoresFromRecord(r, cR["sets"]+1),
	}
}

// ToRecord returns a csv line in this case a []string.
// It is the opposite of NewFromRecord.
func (g Game) ToRecord() []string {
	r := []string{
		g.Country,
		g.Name,
		g.Serie,
		g.Division,
		g.Date,
		"",
		"",
		g.Round,
		g.Category,
		strconv.Itoa(int(g.Teams)),
		"",
		"",
		g.Players[0],
		g.Players[1],
		g.Players[2],
		g.Players[3],
		g.Players[4],
		g.Players[5],
		g.Players[6],
		g.Players[7],
		strconv.Itoa(int(g.Sets)),
	}
	r = append(r, ScoresToRecord(g.Scores)...)
	return r
}

// Tournament returns the tournament of the game
func (g Game) Tournament() Tournament {
	return Tournament{g.Country, g.Name, g.Serie, g.Division}
}
