package model

import (
	"sort"
	"strconv"

	"golang.org/x/exp/slices"
)

// Games is a collections of games and methods to manipulate the collection
type Games []Game

// Len returns the number of games
func (gs Games) Len() int { return len(gs) }

// Swap swaps two games
func (gs Games) Swap(i, j int) { gs[i], gs[j] = gs[j], gs[i] }

// Less returns true if game i is less than game j
func (gs Games) Less(i, j int) bool {
	return gs[i].Compare(gs[j]) < 0
}

func ScoresFromRecord(r []string, start int) []int8 {
	scores := make([]int8, 0)
	for i := start; i < len(r); i++ {
		if r[i] == "" {
			break
		}
		s, err := strconv.Atoi(r[i])
		if err != nil {
			panic(err)
		}
		scores = append(scores, int8(s))
	}
	if len(scores)%2 != 0 {
		panic("Scores must be even")
	}
	return scores
}

func ScoresToRecord(scores []int8) []string {
	r := make([]string, len(scores))
	for i, s := range scores {
		r[i] = strconv.Itoa(int(s))
	}
	return r
}

// Add adds a game to the collection and sorts the collection if sorted is true
func (gs *Games) Add(g Game, sorted bool) {
	copy := *gs
	*gs = append(copy, g)
	if sorted {
		sort.Sort(Games(*gs))
	}
}

// AddBulk adds a collection of games to the collection
func (gs *Games) AddBulk(games []Game) {
	copy := *gs
	*gs = append(copy, games[0:]...)
	sort.Sort(Games(*gs))
}

// Update updates a game in the collection
func (gs *Games) Update(index int, g Game) {
	if len(*gs) == 0 || index < 0 || index >= len(*gs) {
		return
	}
	(*gs)[index] = g
	sort.Sort(Games(*gs))
}

// AddFromRecords reads csv records and creates a collection of games wihout duplicates
func (gs Games) AddFromRecords(records [][]string) Games {
	for _, r := range records {
		gs.Add(NewFromRecord(r), false)
	}
	sort.Sort(Games(gs))
	return gs
}

// Tournaments returns a list of tournaments without duplicates
func (gs Games) Tournaments() []Tournament {
	ts := []Tournament{}
	for _, g := range gs {
		//if _, b := slices.BinarySearch(ts, g.Tournament()); !b {
		if !slices.Contains(ts, g.Tournament()) {
			ts = append(ts, g.Tournament())
		}
	}
	return ts
}

// Dates returns a list without duplicates of all the available dates in Games
func (gs Games) Dates() []string {
	m := map[string]bool{}
	for _, g := range gs {
		m[g.Date] = true
	}
	return mapToList(m)
}

// Rounds returns a list without duplicates of all the rounds in Games
func (gs Games) Rounds() []string {
	m := map[string]bool{}
	for _, g := range gs {
		m[g.Round] = true
	}
	return mapToList(m)
}

// Categories return a list without duplicates of all the categories in Games
func (gs Games) Categories() []string {
	m := map[string]bool{}
	for _, g := range gs {
		m[g.Category] = true
	}
	return mapToList(m)
}

// Lastnames returns a list without duplicates of all the countries in Games
func (gs Games) Lastnames() []string {
	m := map[string]bool{}
	for _, g := range gs {
		for i, p := range g.Players {
			if i%2 == 0 {
				m[p] = true
			}
		}
	}
	return mapToList(m)
}

// Firstnames returns a list without duplicates of all the last names in Games
func (gs Games) Firstnames() []string {
	m := map[string]bool{}
	for _, g := range gs {
		for i, p := range g.Players {
			if i%2 == 1 {
				m[p] = true
			}
		}
	}
	return mapToList(m)
}
