package model

type Game struct {
	Country string
	Name    string
}

func (a Game) Compare(b Game) int {
	// equals
	if a.Country == b.Country &&
		a.Name == b.Name {
		return 0
	}
	// less than
	if a.Country < b.Country {
		return -1
	}
	if a.Country == b.Country &&
		a.Name < b.Name {
		return -1
	}
	// greater than
	return +1
}
