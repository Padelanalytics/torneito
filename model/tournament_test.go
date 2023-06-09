package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var t1 = Tournament{Country: "GERMANY", Name: "Berlin"}
var t2 = Tournament{Country: "GERMANY", Name: "Muenchen"}
var t3 = Tournament{Country: "GERMANY", Name: "Berlin"}
var t4 = Tournament{Country: "SPAIN", Name: "Berlin"}

func TestCompareTournament(t *testing.T) {
	assert.Equal(t, 0, t1.Compare(t3))
	assert.Equal(t, -1, t1.Compare(t2))
	assert.Equal(t, +1, t2.Compare(t1))
	assert.Equal(t, -1, t1.Compare(t4))
}
