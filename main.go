package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/qbit/pgenv"
)

type test struct {
	Tester      string
	ArchName    string
	BootMethod  string
	InstMethod  bool
	Status      bool
	Snapdate    time.Time
	MachineName string
	X11         bool
	Note        string
	Date        time.Time
}

func formatDate(t time.Time) string {
	return t.Format(time.RFC1123)
}

type tests []test

func getTests() (tests, error) {
	var query = `
select
 tester,
 arch.name as ArchName,
 install.name as bootmethod,
 upgrade as instmethod,
 status,
 snap,
 test.name as machinename,
 x11,
 note,
 test.created as date
from test
 join install on install.id = test.installid
 join arch on arch.id = install.archid
`
	var tests tests

	var s = pgenv.ConnStr{}
	s.SetDefaults()

	db, err := sql.Open("postgres", s.ToString())
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	for rows.Next() {
		var t = test{}
		err := rows.Scan(
			&t.Tester,
			&t.ArchName,
			&t.BootMethod,
			&t.InstMethod,
			&t.Status,
			&t.Snapdate,
			&t.MachineName,
			&t.X11,
			&t.Note,
			&t.Date,
		)
		if err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}

	return tests, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"formatDate": formatDate,
	}

	templ, err := template.New("index").Funcs(funcMap).ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	table, err := getTests()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	templ.ExecuteTemplate(w, "index.html", table)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
