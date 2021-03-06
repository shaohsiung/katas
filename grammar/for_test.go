package grammar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFor(t *testing.T) {
	expected := "players: Luca, Allen, Curry, James"
	assert.Equal(t, expected, CommonFor())
	assert.Equal(t, expected, WhileLikeFor())
	assert.Equal(t, expected, InfiniteFor())
	assert.Equal(t, expected, RangeFor())
	assert.Equal(t, expected, RangeMapFor())
}
