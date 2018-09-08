package route

import (
	"database/sql"
	"errors"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nurlansu/go-chuck/models/database"
)

var (
	t = template.Must(template.ParseFiles("public/templates/index.tmpl"))
)

// Handler for "/favicon.ico" route
func faviconHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "public/static/favicon.ico")
	}
}

// Handler for "/" (home) route
func mainHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if fact, err := getRandomFact(db); err == nil {
			t.ExecuteTemplate(w, "index.tmpl", fact)
		} else {
			http.Error(w, err.Error(), 500)
		}
	}
}

// Handler for "/fact/:id" route
func idHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		if fact, err := getTheFact(db, id); err == nil {
			t.ExecuteTemplate(w, "index.tmpl", fact)
		} else {
			http.Error(w, err.Error(), 404)
		}
	}
}

// Handler for "/api" and "/api/:id" route
func apiHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var res Response
		var err error
		id := p.ByName("id")

		if id == "" {
			res.Value, err = getRandomFact(db)
		} else {
			res.Value, err = getTheFact(db, id)
		}

		if err != nil {
			res.Error = err.Error()
		}

		// Marshalling json reponse
		resp := res.Serialize()

		// Setting response headers and writing the response
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}

// getRandomFact returns random Fact from database
func getRandomFact(db *sql.DB) (database.Item, error) {
	return database.GetItem(db, rand.Intn(539))
}

// getTheFact returns Fact specified in 's' from database
func getTheFact(db *sql.DB, s string) (database.Item, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return database.Item{}, err
	} else if id < 0 || id > 540 {
		return database.Item{}, errors.New("no fact with this id :(")
	}

	return database.GetItem(db, id)
}
