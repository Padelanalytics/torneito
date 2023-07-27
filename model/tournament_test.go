package model_test

import (
	"testing"

	m "github.com/paconte/torneito/model"
	"github.com/stretchr/testify/assert"
)

var t1 = m.Tournament{Country: "GERMANY", Name: "Berlin Open", Serie: "GPT-1000", Division: "MO"}
var t2 = m.Tournament{Country: "GERMANY", Name: "Muenchen Open", Serie: "GPT-1000", Division: "MO"}
var t3 = m.Tournament{Country: "GERMANY", Name: "Berlin Open", Serie: "GPT-1000", Division: "MO"}
var t4 = m.Tournament{Country: "SPAIN", Name: "Berlin Open", Serie: "GPT-1000", Division: "MO"}

var games = m.Games{
	m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPT-1000", Division: "MO"},
	m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPT-1000", Division: "WO"},

	m.Game{Country: "GERMANY", Name: "Muenchen Open", Serie: "GPT-1000", Division: "MO"},
	m.Game{Country: "GERMANY", Name: "Muenchen Open", Serie: "GPT-1000", Division: "WO"},

	m.Game{Country: "GERMANY", Name: "Koeln Open", Serie: "GPT-500", Division: "MO"},
	m.Game{Country: "GERMANY", Name: "Koeln Open", Serie: "GPT-500", Division: "WO"},

	m.Game{Country: "WPT", Name: "Madrid Masters", Serie: "WPT-OPEN", Division: "MO"},
	m.Game{Country: "WPT", Name: "Madrid Masters", Serie: "WPT-OPEN", Division: "WO"},
}

func TestTCompare(t *testing.T) {
	assert.Equal(t, 0, t1.Compare(t3))
	assert.Equal(t, -1, t1.Compare(t2))
	assert.Equal(t, +1, t2.Compare(t1))
	assert.Equal(t, -1, t1.Compare(t4))
}

func TestTAdd(t *testing.T) {
	ts := m.Tournaments{}
	assert.Equal(t, 0, len(ts))
	ts.Add(t1)
	assert.Equal(t, 1, len(ts))
	ts.Add(t2)
	assert.Equal(t, 2, len(ts))
	ts.Add(t3)
	assert.Equal(t, 2, len(ts))
	ts.Add(t4)
	assert.Equal(t, 3, len(ts))
}

func TestTRemove(t *testing.T) {
	ts := m.Tournaments{}
	ts.Add(t1)
	ts.Add(t2)
	ts.Add(t3)
	ts.Add(t4)

	assert.Equal(t, 3, len(ts))
	ts.Remove(len(ts) - 1)
	assert.Equal(t, 2, len(ts))
	ts.Remove(1)
	assert.Equal(t, 1, len(ts))
	ts.Remove(0)
	assert.Equal(t, 0, len(ts))
	ts.Remove(0)
	assert.Equal(t, 0, len(ts))
}

func TestFromGame(t *testing.T) {
	ts := m.Tournaments{}
	ts = ts.FromGame(g1)
	assert.Equal(t, 1, len(ts))
	assert.Equal(t, "Berlin Open", ts[0].Name)
}

func TestFromGames(t *testing.T) {
	ts := m.Tournaments{}
	games := m.Games{g1, g2, g3, g4, g5}
	ts = ts.FromGames(games)
	assert.Equal(t, 3, len(ts))
	assert.Equal(t, "Berlin Open", ts[0].Name)
	assert.Equal(t, "Muenchen Open", ts[1].Name)
}

func TestTNames(t *testing.T) {
	ts := m.Tournaments{}
	ts = ts.FromGames(games)
	assert.Equal(t, 4, len(ts.Names()))
}

func TestTCountries(t *testing.T) {
	ts := m.Tournaments{}
	ts = ts.FromGames(games)
	assert.Equal(t, 2, len(ts.Countries()))
}

func TestTDivisions(t *testing.T) {
	ts := m.Tournaments{}
	ts = ts.FromGames(games)
	assert.Equal(t, 2, len(ts.Divisions()))
}

func TestTSeries(t *testing.T) {
	ts := m.Tournaments{}
	ts = ts.FromGames(games)
	assert.Equal(t, 3, len(ts.Series()))
}
