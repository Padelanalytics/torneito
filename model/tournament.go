package model

// Tournament represents a tournament
type Tournament struct {
	Country  string
	Name     string
	Serie    string
	Division string
}

// Compare returns -1 if tournament is less than other,
func (t Tournament) Compare(other Tournament) int {
	// equals
	if t.Country == other.Country &&
		t.Name == other.Name &&
		t.Serie == other.Serie &&
		t.Division == other.Division {
		return 0
	}
	// less than
	if t.Country < other.Country ||
		t.Country == other.Country && t.Name < other.Name ||
		t.Country == other.Country && t.Name == other.Name && t.Serie < other.Serie ||
		t.Country == other.Country && t.Name == other.Name && t.Serie == other.Serie && t.Division < other.Division {
		return -1
	}
	// greater than
	return +1
}
