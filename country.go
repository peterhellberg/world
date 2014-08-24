package world

// Country represents a country
type Country struct {
	Name   string // country name e.g. Austria
	Slug   string // slug based on country name
	Key    string // key (iso alpha2 if available - otherwise alpha3)
	Code   string // country code
	Motor  string // int'l vehicle registration code
	Alpha2 string // iso alpha2 code e.g. AU
	Alpha3 string // iso alpha3 code e.g. AUT
	FIFA   string // FIFA (football) code
	Net    string // internet top level domain (tld) e.g. at
	Region string // region name
	Pop    int    // population of the country
	Area   int    // area of the country
	UN     bool   // un (united nations) flag member (true|false)
	EU     bool   // european union flag member (true|false)
	EURO   bool   // euro (currency) flag member (true|false)
}

type Territory Country
type Supranational Country

func (c *Country) IsISO() bool {
	return c.Alpha3 != ""
}

func (c *Country) IsFIFA() bool {
	return c.FIFA != ""
}
