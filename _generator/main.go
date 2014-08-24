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

	for rows.Next() {
		var (
			name string
			slug string
			key  string
			//place_id     int
			//code         string
			//alt_names    string
			//pop          int
			//area         int
			//continent_id int
			//country_id   int
			//s            bool
			//c            bool
			//d            bool
			//motor        string
			//iso2         string
			iso3  string
			fifa  string
			net   string
			cName string
			//wikipedia    string
		)

		rows.Scan(
			&name,
			&slug,
			&key,
			&iso3,
			&fifa,
			&net,
			&cName)

		countryTemplate.Execute(os.Stdout, &CountryData{
			Name:          name,
			Slug:          slug,
			Key:           key,
			Alpha3:        iso3,
			FIFA:          fifa,
			Net:           net,
			ContinentName: cName,
		})
	}
	rows.Close()
}

type CountryData struct {
	Name          string
	Slug          string
	Key           string
	Alpha3        string
	FIFA          string
	Net           string
	ContinentName string
}

var countriesFileHeader = `// Automatically generated file. Do not edit!

package world
`

var selectCountriesSQL = `SELECT
C.name,
C.slug,
C.key,
C.iso3,
C.fifa,
C.net,
T.Name
FROM COUNTRIES AS C
JOIN CONTINENTS AS T
ON C.continent_id == T.id`

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

var countryTemplate = template.Must(template.New("c").Funcs(funcMap).Parse(`
// {{.Key | ToUpper}} contains the data for {{.Name}}
var {{.Key | ToUpper}} = &Country{
	Name:          "{{.Name}}",
	Key:           "{{.Key}}",
	Alpha3:        "{{.Alpha3}}",
	FIFA:          "{{.FIFA}}",
	Net:           "{{.Net}}",
	ContinentName: "{{.ContinentName}}",
}

World = append(World, {{.Key | ToUpper}})
`))
