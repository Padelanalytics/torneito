package model_test

import (
	"testing"

	m "github.com/paconte/torneito/model"

	"github.com/stretchr/testify/assert"
)

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

func TestAdd(t *testing.T) {

	games := m.Games{}
	games.Add(m.Game{Name: "Berlin", Country: "GER"}, false)
	games.Add(m.Game{Name: "Muenchen", Country: "GER"}, true)
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Muenchen", games[1].Name)

	games.Add(m.Game{Name: "Koeln", Country: "GER"}, true)
	assert.Equal(t, 3, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Koeln", games[1].Name)
	assert.Equal(t, "Muenchen", games[2].Name)

	games.Add(m.Game{Name: "Madrid", Country: "SPAIN"}, true)
	assert.Equal(t, 4, len(games))
	assert.Equal(t, "Madrid", games[3].Name)

	games.Add(m.Game{Name: "Barcelona", Country: "SPAIN"}, true)
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Barcelona", games[3].Name)
	assert.Equal(t, "Madrid", games[4].Name)

}

func TestAddBulk(t *testing.T) {
	games := m.Games{}
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER"},
			{Name: "Muenchen", Country: "GER"},
			{Name: "Koeln", Country: "GER"},
			{Name: "Madrid", Country: "SPAIN"},
			{Name: "Barcelona", Country: "SPAIN"},
		})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Koeln", games[1].Name)
	assert.Equal(t, "Muenchen", games[2].Name)
	assert.Equal(t, "Barcelona", games[3].Name)
	assert.Equal(t, "Madrid", games[4].Name)
}

func TestUpdate(t *testing.T) {
	games := m.Games{}

	// Update an empty collection
	games.Update(0, m.Game{Name: "Berlin", Country: "GER"})
	assert.Equal(t, 0, len(games))

	// Add games
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER"},
			{Name: "Muenchen", Country: "GER"},
			{Name: "Koeln", Country: "GER"},
			{Name: "Madrid", Country: "SPAIN"},
			{Name: "Barcelona", Country: "SPAIN"},
		})

	// Update
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Koeln", games[1].Name)
	assert.Equal(t, "Muenchen", games[2].Name)
	assert.Equal(t, "Barcelona", games[3].Name)
	assert.Equal(t, "Madrid", games[4].Name)

	games.Update(0, m.Game{Name: "Berlin", Country: "xGERMANY"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Berlin", games[4].Name)
	assert.Equal(t, "xGERMANY", games[4].Country)

	games.Update(0, m.Game{Name: "Colonia", Country: "yGER"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Colonia", games[4].Name)
	assert.Equal(t, "yGER", games[4].Country)

	games.Update(0, m.Game{Name: "Munich", Country: "zAlemania"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Munich", games[4].Name)
	assert.Equal(t, "zAlemania", games[4].Country)

	games.Update(4, m.Game{Name: "Dresden", Country: "Alemania"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Dresden", games[0].Name)
	assert.Equal(t, "Alemania", games[0].Country)

	// Out of index
	games.Update(5, m.Game{Name: "Dresden", Country: "Alemania"})
	assert.Equal(t, 5, len(games))
	games.Update(-1, m.Game{Name: "Dresden", Country: "Alemania"})
	assert.Equal(t, 5, len(games))
}

func TestDates(t *testing.T) {
	games := m.Games{}
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER", Date: "2015-01-01"},
			{Name: "Muenchen", Country: "GER", Date: "2015-01-02"},
			{Name: "Koeln", Country: "GER", Date: "2015-01-03"},
			{Name: "Madrid", Country: "SPAIN", Date: "2015-01-04"},
			{Name: "Barcelona", Country: "SPAIN", Date: "2015-01-05"},
			{Name: "Barcelona", Country: "SPAIN", Date: "2015-01-05"},
			{Name: "Madrid", Country: "SPAIN", Date: "2015-01-01"},
		})

	assert.Equal(t, 5, len(games.Dates()))
	for _, s := range []string{"2015-01-01", "2015-01-02", "2015-01-03", "2015-01-04", "2015-01-05"} {
		assert.Contains(t, games.Dates(), s)
	}
}

func TestRounds(t *testing.T) {
	games := m.Games{}
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER", Round: "1"},
			{Name: "Muenchen", Country: "GER", Round: "2"},
			{Name: "Koeln", Country: "GER", Round: "3"},
			{Name: "Madrid", Country: "SPAIN", Round: "4"},
			{Name: "Barcelona", Country: "SPAIN", Round: "5"},
			{Name: "Barcelona", Country: "SPAIN", Round: "5"},
			{Name: "Madrid", Country: "SPAIN", Round: "1"},
		})

	assert.Equal(t, 5, len(games.Rounds()))
	for _, s := range []string{"1", "2", "3", "4", "5"} {
		assert.Contains(t, games.Rounds(), s)
	}
}

func TestCategories(t *testing.T) {
	games := m.Games{}
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER", Category: "1"},
			{Name: "Muenchen", Country: "GER", Category: "2"},
			{Name: "Koeln", Country: "GER", Category: "3"},
			{Name: "Madrid", Country: "SPAIN", Category: "4"},
			{Name: "Barcelona", Country: "SPAIN", Category: "5"},
			{Name: "Barcelona", Country: "SPAIN", Category: "5"},
			{Name: "Madrid", Country: "SPAIN", Category: "1"},
		})

	assert.Equal(t, 5, len(games.Categories()))
	for _, s := range []string{"1", "2", "3", "4", "5"} {
		assert.Contains(t, games.Categories(), s)
	}
}

func TestFirstAndLastNames(t *testing.T) {
	games := m.Games{}
	games.AddBulk(
		[]m.Game{
			{Name: "Berlin", Country: "GER", Players: []string{"1", "2", "3", "4", "5", "6"}},
			{Name: "Muenchen", Country: "GER", Players: []string{"1", "2", "3", "4", "5", "6"}},
			{Name: "Koeln", Country: "GER", Players: []string{"1", "2", "3", "4", "5", "6"}},
			{Name: "Madrid", Country: "SPAIN", Players: []string{"1", "2", "3", "4", "5", "6"}},
			{Name: "Barcelona", Country: "SPAIN", Players: []string{"1", "2", "3", "4", "5", "6"}},
		})

	assert.Equal(t, 3, len(games.Firstnames()))
	assert.Equal(t, 3, len(games.Lastnames()))

	for _, s := range []string{"2", "4", "6"} {
		assert.Contains(t, games.Firstnames(), s)
		assert.NotContains(t, games.Lastnames(), s)
	}
	for _, s := range []string{"1", "3", "5"} {
		assert.Contains(t, games.Lastnames(), s)
		assert.NotContains(t, games.Firstnames(), s)
	}
}

func TestAddFromRecords(t *testing.T) {
	games := m.Games{}
	games = games.AddFromRecords(records)
	assert.Equal(t, 5, len(games))
}
