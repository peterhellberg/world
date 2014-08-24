package world

import "strings"

var (
	World             []*Country
	WorldUN           []*Country
	WorldISO          []*Country
	WorldFIFA         []*Country
	WorldG8           []*Country
	WorldG20          []*Country
	WorldCommonwealth []*Country
	Europe            []*Country
	EuropeEU          []*Country
	EuropeEuro        []*Country
)

type Country struct {
	Name          string // country name e.g. Austria
	Key           string // key (iso alpha2 if available - otherwise alpha3)
	Alpha3        string // iso alpha3 code e.g. AUT
	FIFA          string // FIFA (football) code
	Net           string // internet top level domain (tld) e.g. at
	ContinentName string // continent name
	Kind          string // CTRY|DEPY|SUPR (country|dependency|supranational)
	UN            bool   // un (united nations) flag member (true|false)
	EU            bool   // european union flag member (true|false)
	EURO          bool   // euro (currency) flag member (true|false)
}

func (c *Country) IsISO() bool {
	return c.Alpha3 != ""
}

func (c *Country) IsFIFA() bool {
	return c.FIFA != ""
}

func (c *Country) Slug() string {
	return slugify(c.Name)
}

func slugify(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
}
