package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountries(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(SE.Name, "Sweden")
	assert.Equal(ZA.Slug, "south-africa")
	assert.Equal(AM.Key, "am")
	assert.Equal(AZ.Code, "AZE")
	assert.Equal(BG.Motor, "BG")
	assert.Equal(BO.Alpha3, "BOL")
	assert.Equal(CA.FIFA, "CAN")
	assert.Equal(CU.Net, "cu")
	assert.Equal(DJ.Region, "Africa")
	assert.Equal(EE.Pop, 1340194)
	assert.Equal(FI.Area, 338145)
}

func TestTerritories(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(AX.Name, "Åland Islands")
}

func TestSupranationals(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(EU.Name, "European Union")
}
