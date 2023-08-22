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

var records = [][]string{
	{"WPT", "Adeslas Madrid Open 2021 - Masculino", "WPT-OPEN", "MO", "11.04.2021", "", "", "KO32", "Preprevia", "2", "", "", "Valdés González", "Javier", "Vasquez", "Simon", "Bye", "Bye", "Bye", "Bye", "2", "6", "0", "6", "0"},
	{"WPT", "Adeslas Madrid Open 2021 - Masculino", "WPT-OPEN", "MO", "11.04.2021", "", "", "KO32", "Preprevia", "2", "", "", "Cerezo Casado", "Mario", "Palasi Lozano", "Javier", "Solla", "David Antolín", "Prado Prego", "Manuel", "2", "7", "6", "7", "6"},
	{"WPT", "Adeslas Madrid Open 2021 - Masculino", "WPT-OPEN", "MO", "11.04.2021", "", "", "KO32", "Preprevia", "2", "", "", "Knutsson", "Carl", "Windahl", "Daniel", "Medina Murphy", "Christian", "Gama González", "Juan Carlos", "2", "3", "6", "4", "6"},
	{"WPT", "Adeslas Madrid Open 2021 - Masculino", "WPT-OPEN", "MO", "11.04.2021", "", "", "KO32", "Preprevia", "2", "", "", "Cremona", "Simone", "Cattaneo", "Daniele", "Muñoz Enrile", "Jaime", "García Mora", "Javier", "2", "4", "6", "4", "6"},
	{"WPT", "Adeslas Madrid Open 2021 - Masculino", "WPT-OPEN", "MO", "11.04.2021", "", "", "KO32", "Preprevia", "2", "", "", "Muñoz Baixas", "Jordi", "Aliaga Cruz", "José", "Molina Domínguez", "Alberto", "Antonelly Carballo", "Ionjan", "2", "6", "3", "6", "4"},
}

func TestScoresFromRecord(t *testing.T) {
	scores := m.ScoresFromRecord([]string{"6", "0", "6", "0"}, 0)
	assert.Equal(t, []int8{6, 0, 6, 0}, scores)

	scores = m.ScoresFromRecord([]string{"6", "0", "6", "0", "6", "0", "6", "0", "6", "0"}, 0)
	assert.Equal(t, []int8{6, 0, 6, 0, 6, 0, 6, 0, 6, 0}, scores)

	scores = m.ScoresFromRecord([]string{"6", "0", ""}, 0)
	assert.Equal(t, []int8{6, 0}, scores)
}

func TestScoresFromRecordPanicOdd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m.ScoresFromRecord([]string{"6", "0", "6"}, 0)
}

func TestScoresFromRecordPanicAtoi(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m.ScoresFromRecord([]string{"6", "0", "a", "0"}, 0)
}

func TestToRecords(t *testing.T) {
	games := []m.Game{}
	newRecords := make([][]string, len(records))
	for _, r := range records {
		games = append(games, m.NewFromRecord(r))
	}
	assert.Equal(t, len(records), len(games))

	for i, g := range games {
		newRecords[i] = g.ToRecord()
	}
	assert.Equal(t, len(records), len(newRecords))
	assert.Equal(t, records, newRecords)
}
