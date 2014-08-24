package world

var (
	// Countries is a map of all countries in the world
	Countries = make(map[string]*Country)

	// Territories is a map of all territories
	Territories = make(map[string]*Territory)

	// Supranationals is a map of all supranationals
	Supranationals = make(map[string]*Supranational)

	// Europe is a map of all countries in Europe
	Europe = make(map[string]*Country)
)
