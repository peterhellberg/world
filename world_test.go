package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(len(World), 245)
}
