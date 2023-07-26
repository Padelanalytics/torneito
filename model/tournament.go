package model

type Tournament struct {
	Country  string
	Name     string
	Serie    string
	Division string
}

func (a Tournament) Compare(b Tournament) int {
	// equals
	if a.Country == b.Country &&
		a.Name == b.Name &&
		a.Serie == b.Serie &&
		a.Division == b.Division {
		return 0
	}
	// less than
	if a.Country < b.Country ||
		a.Country == b.Country && a.Name < b.Name ||
		a.Country == b.Country && a.Name == b.Name && a.Serie < b.Serie ||
		a.Country == b.Country && a.Name == b.Name && a.Serie == b.Serie && a.Division < b.Division {
		return -1
	}
	// greater than
	return +1
}
