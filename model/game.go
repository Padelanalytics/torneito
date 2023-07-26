package model

type Game struct {
	Country  string
	Name     string
	Serie    string
	Division string
	Date     string
	Round    string
	Category string
	Teams    string
	Players  []string
	Sets     uint8
	Local    []int8
	Visitor  []int8
}

func (a Game) Compare(b Game) int {
	// equals
	if a.Country == b.Country &&
		a.Name == b.Name &&
		a.Serie == b.Serie &&
		a.Division == b.Division &&
		a.Category == b.Category &&
		a.Round == b.Round {
		return 0
	}
	// less than
	if a.Country < b.Country ||
		a.Country == b.Country && a.Name < b.Name ||
		a.Country == b.Country && a.Name == b.Name && a.Serie < b.Serie ||
		a.Country == b.Country && a.Name == b.Name && a.Serie == b.Serie && a.Division < b.Division ||
		a.Country == b.Country && a.Name == b.Name && a.Serie == b.Serie && a.Division == b.Division && cCategory[a.Category] < cCategory[b.Category] ||
		a.Country == b.Country && a.Name == b.Name && a.Serie == b.Serie && a.Division == b.Division && cCategory[a.Category] < cCategory[b.Category] && cRound[a.Round] < cRound[b.Round] {
		return -1
	}
	// greater than
	return +1
}

func (a Game) IsTournament(t Tournament) bool {
	return a.Country == t.Country &&
		a.Name == t.Name &&
		a.Serie == t.Serie &&
		a.Division == t.Division
}
