package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var date = "10.20.2020"
var g1 = Game{Country: "GERMANY", Name: "Berlin", Date: date}
var g2 = Game{Country: "GERMANY", Name: "Muenchen", Date: date}
var g3 = Game{Country: "GERMANY", Name: "Berlin", Date: date}
var g4 = Game{Country: "SPAIN", Name: "Berlin", Date: date}
var g5 = Game{Country: "GERMANY", Name: "Berlin", Date: "20.10.2021"}

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, g1.Compare(g3))
	assert.Equal(t, -1, g1.Compare(g2))
	assert.Equal(t, +1, g2.Compare(g1))
	assert.Equal(t, -1, g1.Compare(g4))
	assert.Equal(t, -1, g1.Compare(g5))
}
