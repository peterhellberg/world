package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountry(t *testing.T) {
	assert := assert.New(t)

	c := Country{Name: "Sweden"}

	assert.Equal(c.Name, "Sweden")
}
