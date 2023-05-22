package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	games := Games{}

	games = games.Add(Game{Name: "Berlin", Country: "GER"})
	games = games.Add(Game{Name: "Muenchen", Country: "GER"})
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Muenchen", games[1].Name)

	games = games.Add(Game{Name: "Koeln", Country: "GER"})
	assert.Equal(t, 3, len(games))
	assert.Equal(t, "Berlin", games[0].Name)
	assert.Equal(t, "Koeln", games[1].Name)
	assert.Equal(t, "Muenchen", games[2].Name)

	games = games.Add(Game{Name: "Madrid", Country: "SPAIN"})
	assert.Equal(t, 4, len(games))
	assert.Equal(t, "Madrid", games[3].Name)

	games = games.Add(Game{Name: "Barcelona", Country: "SPAIN"})
	assert.Equal(t, 5, len(games))
	assert.Equal(t, "Barcelona", games[3].Name)
	assert.Equal(t, "Madrid", games[4].Name)
}

func TestRemove(t *testing.T) {
	games := Games{}

	games = games.Add(Game{Name: "Berlin", Country: "GER"})
	games = games.Add(Game{Name: "Muenchen", Country: "GER"})
	games = games.Add(Game{Name: "Koeln", Country: "GER"})
	games = games.Add(Game{Name: "Madrid", Country: "SPAIN"})
	games = games.Add(Game{Name: "Barcelona", Country: "SPAIN"})

	// remove first
	games = games.Remove(0)
	assert.Equal(t, 4, len(games))
	assert.Equal(t, "Koeln", games[0].Name)

	// remove last
	games = games.Remove(3)
	assert.Equal(t, 3, len(games))
	assert.Equal(t, "Barcelona", games[2].Name)

	// remove middle
	games = games.Remove(1)
	assert.Equal(t, 2, len(games))
	assert.Equal(t, "Koeln", games[0].Name)
	assert.Equal(t, "Barcelona", games[1].Name)

	// out of index
	games = games.Remove(2)
	assert.Equal(t, 2, len(games))
	games = games.Remove(20)
	assert.Equal(t, 2, len(games))
	games = games.Remove(-1)
	assert.Equal(t, 2, len(games))
}
func TestAddBulk(t *testing.T) {
	games := Games{}
	games = games.AddBulk(
		[]Game{
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
