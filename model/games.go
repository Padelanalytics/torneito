package model

import (
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

func (a Games) Remove(index int) Games {
	if index < 0 || index >= len(a) {
		return a
	}
	return append(a[:index], a[index+1:]...)
}
