package model

import (
	"fmt"
	"sort"
)

type Games []Game

func (a Games) Len() int { return len(a) }

func (a Games) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Games) Less(i, j int) bool {
	return a[i].Compare(a[j]) < 0
}

func (a Games) Add(g Game) Games {
	a = append(a, g)
	sort.Sort(Games(a))
	return a
}

func (a Games) AddBulk(gs []Game) Games {
	a = append(a, gs[0:]...)
	sort.Sort(Games(a))
	return a
}

func (a Games) Remove(index int) Games {
	if index < 0 || index >= len(a) {
		return a
	}
	return append(a[:index], a[index+1:]...)
}

func (a Games) FromRecords(records [][]string) Games {
	for _, r := range records {
		fmt.Println(r)
		a = a.Add(Game{Country: r[0], Name: r[1]})
	}
	return a
}
