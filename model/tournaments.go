package model

type Tournaments []Tournament

func (ts Tournaments) Len() int { return len(ts) }

func (ts Tournaments) Swap(i, j int) { ts[i], ts[j] = ts[j], ts[i] }

func (ts Tournaments) Less(i, j int) bool {
	return ts[i].Compare(ts[j]) < 0
}

func (ts Tournaments) contains(s []Tournament, e Tournament) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (ts *Tournaments) Add(t Tournament) {
	if !ts.contains(*ts, t) {
		copy := *ts
		*ts = append(copy, t)
	}
}

func (ts *Tournaments) Remove(index int) {
	if index < 0 || index >= len(*ts) {
		return
	}
	copy := *ts
	*ts = append(copy[:index], copy[index+1:]...)
}

func (ts Tournaments) FromGame(g Game) Tournaments {
	ts.Add(Tournament{g.Country, g.Name, g.Serie, g.Division})
	return ts
}

func (ts Tournaments) FromGames(games Games) Tournaments {
	for _, g := range games {
		ts = ts.FromGame(g)
	}
	return ts
}

func (ts Tournaments) Names() []string {
	ns := map[string]bool{}
	for _, t := range ts {
		ns[t.Name] = true
	}
	return mapToList(ns)
}

func (ts Tournaments) Countries() []string {
	cs := map[string]bool{}
	for _, t := range ts {
		cs[t.Country] = true
	}
	return mapToList(cs)
}

func (ts Tournaments) Series() []string {
	ss := map[string]bool{}
	for _, t := range ts {
		ss[t.Serie] = true
	}
	return mapToList(ss)
}

func (ts Tournaments) Divisions() []string {
	ds := map[string]bool{}
	for _, t := range ts {
		ds[t.Division] = true
	}
	return mapToList(ds)
}

func mapToList(m map[string]bool) []string {
	l := []string{}
	for k := range m {
		l = append(l, k)
	}
	return l
}
