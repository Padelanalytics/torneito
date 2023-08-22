package model_test

import (
	"testing"

	m "github.com/paconte/torneito/model"
	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	ts := games.Tournaments()
	assert.Equal(t, 4, len(m.Names(ts)))
}

func TestCountries(t *testing.T) {
	ts := games.Tournaments()
	assert.Equal(t, 2, len(m.Countries(ts)))
}

func TestDivisions(t *testing.T) {
	ts := games.Tournaments()
	assert.Equal(t, 2, len(m.Divisions(ts)))
}

func TestSeries(t *testing.T) {
	ts := games.Tournaments()
	assert.Equal(t, 3, len(m.Series(ts)))
}
