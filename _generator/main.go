package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var dbFile = flag.String("db", "world.db", "Path to the world.db database")

func main() {
	flag.Parse()

	// Look for the database
	if _, err := os.Stat(*dbFile); os.IsNotExist(err) {
		fmt.Printf("Missing database file: %s\n", *dbFile)
		return
	}

	// Open the database
	db, err := sql.Open("sqlite3", *dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	generateCountries(db, os.Stdout)
}

func generateCountries(db *sql.DB, f *os.File) {
	f.WriteString(countriesFileHeader)

	rows, err := db.Query(selectCountriesSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	worldCountries := []string{}
	europeCountries := []string{}
	territories := []string{}
	supras := []string{}

	for rows.Next() {
		var (
			name      = ""
			slug      = ""
			key       = ""
			code      = ""
			motor     = ""
			iso3      = ""
			fifa      = ""
			net       = ""
			cName     = ""
			pop       = int(-1)
			area      = int(-1)
			country   = false
			territory = false
			supra     = false
		)

		rows.Scan(
			&name,
			&slug,
			&key,
			&code,
			&iso3,
			&fifa,
			&net,
			&cName,
			&pop,
			&area,
			&supra,
			&country,
			&territory,
			&motor,
		)

		cd := &CountryData{
			Name:   name,
			Slug:   slug,
			Key:    key,
			Code:   code,
			Motor:  motor,
			Alpha3: iso3,
			FIFA:   fifa,
			Net:    net,
			Region: cName,
			Pop:    pop,
			Area:   area,
		}

		if country {
			worldCountries = append(worldCountries, key)

			if cName == "Europe" {
				europeCountries = append(europeCountries, key)
			}

			countryTemplate.Execute(os.Stdout, cd)
		} else if supra {
			supras = append(supras, key)
			supraTemplate.Execute(os.Stdout, cd)
		} else {
			territories = append(territories, key)
			territoryTemplate.Execute(os.Stdout, cd)
		}
	}
	rows.Close()

	f.WriteString("\nfunc init() {")
	f.WriteString("\n\t// Countries in the world")
	for _, key := range worldCountries {
		k := strings.ToUpper(key)
		f.WriteString("\n\tCountries[\"" + k + "\"] = " + k)
	}

	f.WriteString("\n\n\t// Countries in Europe")
	for _, key := range europeCountries {
		k := strings.ToUpper(key)
		f.WriteString("\n\tEurope[\"" + k + "\"] = " + k)
	}

	f.WriteString("\n\n\t// Supranationals")
	for _, key := range supras {
		k := strings.ToUpper(key)
		f.WriteString("\n\tSupranationals[\"" + k + "\"] = " + k)
	}

	f.WriteString("\n\n\t// Territories")
	for _, key := range territories {
		k := strings.ToUpper(key)
		f.WriteString("\n\tTerritories[\"" + k + "\"] = " + k)
	}
	f.WriteString("\n}\n")
}

type CountryData struct {
	Name   string
	Slug   string
	Key    string
	Code   string
	Motor  string
	Alpha3 string
	FIFA   string
	Net    string
	Region string
	Pop    int
	Area   int
}

var countriesFileHeader = `// Automatically generated file. Do not edit!

package world
`

var selectCountriesSQL = `SELECT
C.name,
C.slug,
C.key,
C.code,
C.iso3,
C.fifa,
C.net,
T.Name,
C.pop,
C.area,
C.s,
C.c,
C.d,
C.motor
FROM COUNTRIES AS C
JOIN CONTINENTS AS T
ON C.continent_id == T.id
ORDER BY C.key`

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

var countryTemplate = template.Must(template.New("c").Funcs(funcMap).Parse(`
// {{.Key | ToUpper}} contains data for {{.Name}}
var {{.Key | ToUpper}} = &Country{
	Name:   "{{.Name}}",
	Slug:   "{{.Slug}}",
	Key:    "{{.Key}}",
	Code:   "{{.Code}}",
	Motor:  "{{.Motor}}",
	Alpha3: "{{.Alpha3}}",
	FIFA:   "{{.FIFA}}",
	Net:    "{{.Net}}",
	Region: "{{.Region}}",
	Pop:    {{.Pop}},
	Area:   {{.Area}},
}
`))

var territoryTemplate = template.Must(template.New("t").Funcs(funcMap).Parse(`
// {{.Key | ToUpper}} contains data for the territory {{.Name}}
var {{.Key | ToUpper}} = &Territory{
	Name:   "{{.Name}}",
	Slug:   "{{.Slug}}",
	Key:    "{{.Key}}",
	Code:   "{{.Code}}",
	Motor:  "{{.Motor}}",
	Alpha3: "{{.Alpha3}}",
	FIFA:   "{{.FIFA}}",
	Net:    "{{.Net}}",
	Region: "{{.Region}}",
	Pop:    {{.Pop}},
	Area:   {{.Area}},
}
`))

var supraTemplate = template.Must(template.New("s").Funcs(funcMap).Parse(`
// {{.Key | ToUpper}} contains data for the supranational {{.Name}}
var {{.Key | ToUpper}} = &Supranational{
	Name:   "{{.Name}}",
	Slug:   "{{.Slug}}",
	Key:    "{{.Key}}",
	Code:   "{{.Code}}",
	Motor:  "{{.Motor}}",
	Alpha3: "{{.Alpha3}}",
	FIFA:   "{{.FIFA}}",
	Net:    "{{.Net}}",
	Region: "{{.Region}}",
	Pop:    {{.Pop}},
	Area:   {{.Area}},
}
`))
