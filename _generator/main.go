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

	keys := []string{}

	for rows.Next() {
		var (
			name  = ""
			slug  = ""
			key   = ""
			code  = ""
			motor = ""
			iso3  = ""
			fifa  = ""
			net   = ""
			cName = ""
			pop   = int(-1)
			area  = int(-1)
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
			&motor,
		)

		keys = append(keys, key)

		countryTemplate.Execute(os.Stdout, &CountryData{
			Name:          name,
			Slug:          slug,
			Key:           key,
			Code:          code,
			Motor:         motor,
			Alpha3:        iso3,
			FIFA:          fifa,
			Net:           net,
			ContinentName: cName,
			Pop:           pop,
			Area:          area,
		})
	}
	rows.Close()

	f.WriteString("\nfunc init() {")
	for _, key := range keys {
		k := strings.ToUpper(key)
		f.WriteString("\n\tCountries[\"" + k + "\"] = " + k)
	}
	f.WriteString("\n}")
}

type CountryData struct {
	Name          string
	Slug          string
	Key           string
	Code          string
	Motor         string
	Alpha3        string
	FIFA          string
	Net           string
	ContinentName string
	Pop           int
	Area          int
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
C.motor
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
	Code:          "{{.Code}}",
	Motor:         "{{.Motor}}",
	Alpha3:        "{{.Alpha3}}",
	FIFA:          "{{.FIFA}}",
	Net:           "{{.Net}}",
	ContinentName: "{{.ContinentName}}",
	Pop:           {{.Pop}},
	Area:          {{.Area}},
}
`))
