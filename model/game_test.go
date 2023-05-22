package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	g1 := Game{Country: "GERMANY", Name: "Berlin"}
	g2 := Game{Country: "GERMANY", Name: "Muenchen"}
	g3 := Game{Country: "GERMANY", Name: "Berlin"}
	g4 := Game{Country: "SPAIN", Name: "Berlin"}

	assert.Equal(t, 0, g1.Compare(g3))
	assert.Equal(t, -1, g1.Compare(g2))
	assert.Equal(t, +1, g2.Compare(g1))
	assert.Equal(t, -1, g1.Compare(g4))
}
