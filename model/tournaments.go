package model

// Tournaments is a collections of tournaments and methods to manipulate the collection
type Tournaments []Tournament

// Len returns the number of tournaments
func (ts Tournaments) Len() int { return len(ts) }

// Swap swaps two tournaments
func (ts Tournaments) Swap(i, j int) { ts[i], ts[j] = ts[j], ts[i] }

// Less returns true if tournament i is less than tournament j
func (ts Tournaments) Less(i, j int) bool {
	return ts[i].Compare(ts[j]) < 0
}

// contains returns true if tournament t is in the collection ts
func (ts Tournaments) contains(t Tournament) bool {
	for _, tournie := range ts {
		if tournie == t {
			return true
		}
	}
	return false
}

// Add adds a tournament to the collection if it does not exist
func (ts *Tournaments) Add(t Tournament) {
	if !ts.contains(t) {
		copy := *ts
		*ts = append(copy, t)
	}
}

// Remove removes a tournament from the collection if it exists
func (ts *Tournaments) Remove(index int) {
	if index < 0 || index >= len(*ts) {
		return
	}
	copy := *ts
	*ts = append(copy[:index], copy[index+1:]...)
}

// FromGame adds the extracted tournament from a game to the
// tournaments collection if it does not exist
func (ts Tournaments) FromGame(g Game) Tournaments {
	ts.Add(Tournament{g.Country, g.Name, g.Serie, g.Division})
	return ts
}

// FromGames adds a collection of extracted tournaments from a collection of games
// to the tournaments, every tournament is added if it does not exist
func (ts Tournaments) FromGames(games Games) Tournaments {
	for _, g := range games {
		ts = ts.FromGame(g)
	}
	return ts
}

// Names returns a list of tournament names without duplicates
func (ts Tournaments) Names() []string {
	ns := map[string]bool{}
	for _, t := range ts {
		ns[t.Name] = true
	}
	return mapToList(ns)
}

// Countries returns a list of tournament countries without duplicates
func (ts Tournaments) Countries() []string {
	cs := map[string]bool{}
	for _, t := range ts {
		cs[t.Country] = true
	}
	return mapToList(cs)
}

// Series returns a list of tournament series without duplicates
func (ts Tournaments) Series() []string {
	ss := map[string]bool{}
	for _, t := range ts {
		ss[t.Serie] = true
	}
	return mapToList(ss)
}

// Divisions returns a list of tournament divisions without duplicates
func (ts Tournaments) Divisions() []string {
	ds := map[string]bool{}
	for _, t := range ts {
		ds[t.Division] = true
	}
	return mapToList(ds)
}

// mapToList converts a map to a list
func mapToList(m map[string]bool) []string {
	l := []string{}
	for k := range m {
		l = append(l, k)
	}
	return l
}
