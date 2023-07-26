package model

import (
	"sort"
	"strconv"
)

type Games []Game

func (gs Games) Len() int { return len(gs) }

func (gs Games) Swap(i, j int) { gs[i], gs[j] = gs[j], gs[i] }

func (gs Games) Less(i, j int) bool {
	return gs[i].Compare(gs[j]) < 0
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

func (gs *Games) Add(g Game) {
	copy := *gs
	*gs = append(copy, g)
	sort.Sort(Games(*gs))
}

func (gs *Games) AddBulk(games []Game) {
	copy := *gs
	*gs = append(copy, games[0:]...)
	sort.Sort(Games(*gs))
}

func (gs *Games) Remove(index int) {
	if index < 0 || index >= len(*gs) {
		return
	}
	copy := *gs
	*gs = append(copy[:index], copy[index+1:]...)
}

func (gs *Games) FromRecord(r []string) {
	sets, _ := strconv.Atoi(r[cR["sets"]])
	local, visitor := getResult(r, cR["sets"]+1)
	g := Game{
		Country:  r[cR["country"]],
		Name:     r[cR["name"]],
		Serie:    r[cR["serie"]],
		Division: r[cR["division"]],
		Date:     r[cR["date"]],
		Round:    r[cR["round"]],
		Category: r[cR["category"]],
		Teams:    r[cR["teams"]],
		Players:  []string{r[cR["p1_last"]], r[cR["p1_first"]], r[cR["p2_last"]], r[cR["p2_first"]], r[cR["p3_last"]], r[cR["p3_first"]], r[cR["p4_last"]], r[cR["p4_first"]]},
		Sets:     uint8(sets),
		Local:    local,
		Visitor:  visitor,
	}
	copy := *gs
	*gs = append(copy, g)
}

func (gs Games) FromRecords(records [][]string) Games {
	for _, r := range records {
		gs.FromRecord(r)
	}
	sort.Sort(Games(gs))
	return gs
}
