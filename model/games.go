package model

import (
	"sort"
	"strconv"
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

func getResult(r []string, idx int) ([]int8, []int8) {
	i := idx
	j := 0
	local := make([]int8, 0)
	visitor := make([]int8, 0)
	var score string
	for i < len(r) {
		score = r[i]
		if score == "" {
			break
		}
		if i%2 == 0 { // visitor
			s, err := strconv.Atoi(score)
			if err != nil {
				panic(err)
			}
			visitor = append(visitor, int8(s))
		} else { // local
			s, err := strconv.Atoi(score)
			if err != nil {
				panic(err)
			}
			local = append(local, int8(s))
		}
		j++
		i = idx + j
	}
	return local, visitor
}

// Add adds a game to the collection
func (gs *Games) Add(g Game) {
	copy := *gs
	*gs = append(copy, g)
	sort.Sort(Games(*gs))
}

// AddBulk adds a collection of games to the collection
func (gs *Games) AddBulk(games []Game) {
	copy := *gs
	*gs = append(copy, games[0:]...)
	sort.Sort(Games(*gs))
}

// Remove removes a game from the collection
func (gs *Games) Remove(index int) {
	if index < 0 || index >= len(*gs) {
		return
	}
	copy := *gs
	*gs = append(copy[:index], copy[index+1:]...)
}

// Update updates a game in the collection
func (gs *Games) Update(index int, g Game) {
	if index < 0 || index >= len(*gs) {
		return
	}
	(*gs)[index] = g
	sort.Sort(Games(*gs))
}

// FromRecord reads a csv record and creates a game
func (gs *Games) FromRecord(r []string) {
	sets, _ := strconv.Atoi(r[cR["sets"]])
	teams, _ := strconv.Atoi(r[cR["teams"]])
	g := Game{
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
	copy := *gs
	*gs = append(copy, g)
}

// FromRecords reads a csv records and creates a collection of games wihout duplicates
func (gs Games) FromRecords(records [][]string) Games {
	for _, r := range records {
		gs.FromRecord(r)
	}
	sort.Sort(Games(gs))
	return gs
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
			if i%2 == 1 {
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
			if i%2 != 0 {
				m[p] = true
			}
		}
	}
	return mapToList(m)
}
