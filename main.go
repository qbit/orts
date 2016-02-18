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

type Test struct {
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

type Tests []Test

func getTests() (Tests, error) {
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
	var tests Tests

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
	for rows.Next() {
		var t = Test{}
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
	var templates = template.Must(template.ParseFiles("index.html"))
	p, err := getTests()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	templates.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
