package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountries(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(SE.Name, "Sweden")
	assert.Equal(ZA.Slug, "south-africa")
}
