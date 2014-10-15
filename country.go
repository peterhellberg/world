package world

// Country represents a country
type Country struct {
	Name   string // country name e.g. Austria
	Slug   string // slug based on country name
	Key    string // key (iso alpha2 if available - otherwise alpha3)
	Code   string // country code
	Motor  string // int'l vehicle registration code
	Alpha3 string // iso alpha3 code e.g. AUT
	FIFA   string // FIFA (football) code
	Net    string // internet top level domain (tld) e.g. at
	Region string // region name
	Pop    int    // population of the country
	Area   int    // area of the country
}

// Territory is a part of a country
type Territory Country

// Supranational is a union of countries (EU)
type Supranational Country

// IsISO returns a boolean if the country has a 3 letter alpha code
func (c *Country) IsISO() bool {
	return c.Alpha3 != ""
}

// IsFIFA returns a boolean if the country is part of FIFA
func (c *Country) IsFIFA() bool {
	return c.FIFA != ""
}
