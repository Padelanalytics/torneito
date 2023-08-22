package model

// Names returns a list of tournament names without duplicates
func Names(ts []Tournament) []string {
	ns := map[string]bool{}
	for _, t := range ts {
		ns[t.Name] = true
	}
	return mapToList(ns)
}

// Countries returns a list of tournament countries without duplicates
func Countries(ts []Tournament) []string {
	cs := map[string]bool{}
	for _, t := range ts {
		cs[t.Country] = true
	}
	return mapToList(cs)
}

// Series returns a list of tournament series without duplicates
func Series(ts []Tournament) []string {
	ss := map[string]bool{}
	for _, t := range ts {
		ss[t.Serie] = true
	}
	return mapToList(ss)
}

// Divisions returns a list of tournament divisions without duplicates
func Divisions(ts []Tournament) []string {
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
