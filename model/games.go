package model

import (
	"fmt"
	"sort"
)

type Games []Game

func (gs Games) Len() int { return len(gs) }

func (gs Games) Swap(i, j int) { gs[i], gs[j] = gs[j], gs[i] }

func (gs Games) Less(i, j int) bool {
	return gs[i].Compare(gs[j]) < 0
}

func (gs Games) Add(g Game) Games {
	gs = append(gs, g)
	sort.Sort(Games(gs))
	return gs
}

func (gs Games) AddBulk(games []Game) Games {
	gs = append(gs, games[0:]...)
	sort.Sort(Games(gs))
	return gs
}

func (gs Games) Remove(index int) Games {
	if index < 0 || index >= len(gs) {
		return gs
	}
	return append(gs[:index], gs[index+1:]...)
}

func (gs Games) FromRecords(records [][]string) Games {
	for _, r := range records {
		fmt.Println(r)
		gs = gs.Add(Game{Country: r[0], Name: r[1]})
	}
	return gs
}
