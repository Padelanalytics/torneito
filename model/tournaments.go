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
	ns := []string{}
	for _, t := range ts {
		ns = append(ns, t.Name)
	}
	return ns
}

func (ts Tournaments) Countries() []string {
	cs := []string{}
	for _, t := range ts {
		cs = append(cs, t.Country)
	}
	return cs
}

func (ts Tournaments) Series() []string {
	ss := []string{}
	for _, t := range ts {
		ss = append(ss, t.Serie)
	}
	return ss
}

func (ts Tournaments) Divisions() []string {
	ds := []string{}
	for _, t := range ts {
		ds = append(ds, t.Division)
	}
	return ds
}
