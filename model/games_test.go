package model_test

import (
	"testing"

	m "github.com/paconte/torneito/model"

	"github.com/stretchr/testify/assert"
)

func TestGsAdd(t *testing.T) {

	games := m.Games{}

	games.Add(m.Game{Name: "Berlin", Country: "GER"})
	games.Add(m.Game{Name: "Muenchen", Country: "GER"})
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Muenchen", games[1].Name)

	games.Add(m.Game{Name: "Koeln", Country: "GER"})
	assert.Equal(t, 3, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Koeln", games[1].Name)
	assert.Equal(t, "Muenchen", games[2].Name)

	games.Add(m.Game{Name: "Madrid", Country: "SPAIN"})
	assert.Equal(t, 4, len(games))
	assert.Equal(t, "Madrid", games[3].Name)

	games.Add(m.Game{Name: "Barcelona", Country: "SPAIN"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Barcelona", games[3].Name)
	assert.Equal(t, "Madrid", games[4].Name)

}

func TestGsRemove(t *testing.T) {

	games := m.Games{}

	games.Add(m.Game{Name: "Berlin", Country: "GER"})
	games.Add(m.Game{Name: "Muenchen", Country: "GER"})
	games.Add(m.Game{Name: "Koeln", Country: "GER"})
	games.Add(m.Game{Name: "Madrid", Country: "SPAIN"})
	games.Add(m.Game{Name: "Barcelona", Country: "SPAIN"})

	// remove first
	games.Remove(0)
	assert.Equal(t, 4, len(games))
	assert.Equal(t, "Koeln", games[0].Name)

	// remove last
	games.Remove(3)
	assert.Equal(t, 3, len(games))
	assert.Equal(t, "Barcelona", games[2].Name)

	// remove middle
	games.Remove(1)
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Koeln", games[0].Name)
	assert.Equal(t, "Barcelona", games[1].Name)

	// out of index
	games.Remove(2)
	assert.Equal(t, 2, len(games))
	games.Remove(20)
	assert.Equal(t, 2, len(games))
	games.Remove(-1)
	assert.Equal(t, 2, len(games))
}

func TestGsAddBulk(t *testing.T) {
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
