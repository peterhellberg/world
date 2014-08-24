package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {
	assert := assert.New(t)

	// Countries
	assert.Equal(196, len(Countries))

	// Territories
	assert.Equal(48, len(Territories))

	// Supranationals
	assert.Equal(1, len(Supranationals))

	// Europe
	assert.Equal(49, len(Europe))
	assert.Equal(Europe["DE"].Motor, "D")
}
