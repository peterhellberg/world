package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {
	assert := assert.New(t)

	// Countries
	assert.Equal(len(Countries), 245)

	// Europe
	assert.Equal(len(Europe), 55)
	assert.Equal(Europe["DE"].Motor, "D")
}
