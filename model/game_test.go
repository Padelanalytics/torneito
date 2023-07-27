package model_test

import (
	"testing"

	m "github.com/paconte/torneito/model"
	"github.com/stretchr/testify/assert"
)

var date = "10.20.2020"
var g1 = m.Game{Country: "GERMANY", Name: "Berlin Open", Date: date}
var g2 = m.Game{Country: "GERMANY", Name: "Muenchen Open", Date: date}
var g3 = m.Game{Country: "GERMANY", Name: "Berlin Open", Date: date}
var g4 = m.Game{Country: "SPAIN", Name: "Berlin Open", Date: date}
var g5 = m.Game{Country: "GERMANY", Name: "Berlin Open", Date: "20.10.2021"}

func TestGCompare(t *testing.T) {
	assert.Equal(t, 0, g1.Compare(g3))
	assert.Equal(t, -1, g1.Compare(g2))
	assert.Equal(t, +1, g2.Compare(g1))
	assert.Equal(t, -1, g1.Compare(g4))
	assert.Equal(t, -1, g1.Compare(g5))
}

func TestGIsTournament(t *testing.T) {
	game1 := m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPS-1000", Division: "MO", Date: date}
	game2 := m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPS-1000", Division: "WO", Date: date}
	game3 := m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPS-500", Division: "MO", Date: date}
	game4 := m.Game{Country: "GERMANY", Name: "Berlin Open II", Serie: "GPS-1000", Division: "MO", Date: date}
	game5 := m.Game{Country: "SPAIN", Name: "Berlin Open", Serie: "GPS-1000", Division: "MO", Date: date}
	game6 := m.Game{Country: "GERMANY", Name: "Berlin Open", Serie: "GPS-1000", Division: "MO", Date: "20.10.2021"}

	tournament := m.Tournament{Country: "GERMANY", Name: "Berlin Open", Serie: "GPS-1000", Division: "MO"}

	assert.Equal(t, true, game1.IsTournament(tournament))
	assert.Equal(t, false, game2.IsTournament(tournament))
	assert.Equal(t, false, game3.IsTournament(tournament))
	assert.Equal(t, false, game4.IsTournament(tournament))
	assert.Equal(t, false, game5.IsTournament(tournament))
	assert.Equal(t, true, game6.IsTournament(tournament))
}
