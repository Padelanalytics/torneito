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

func (gs Games) FromRecords(records [][]string) Games {
	for _, r := range records {
		fmt.Println(r)
		gs.Add(Game{Country: r[0], Name: r[1], Date: r[4]})
	}
	return gs
}
