package route

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
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

// Handler for "/fact/:id" route
func idHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		if id, err := strconv.Atoi(id); id > 0 && id < 540 && err == nil {
			fact, _ := getTheFact(db, id)
			t.ExecuteTemplate(w, "index.tmpl", fact)
		} else {
			http.Error(w, "404 page not found", 404)
		}
	}
}

// Handler for "/" (home) route
func mainHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fact, _ := getTheFact(db, 0)
		t.ExecuteTemplate(w, "index.tmpl", fact)
	}
}

// Handler for "/api" route
func apiHandler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var fact database.Item
		id := p.ByName("id")
		if id, err := strconv.Atoi(id); id > 0 && id < 540 && err == nil {
			fact, _ = getTheFact(db, id)
		} else {
			fact, _ = getTheFact(db, 0)
		}
		res, err := json.Marshal(fact)
		if err != nil {
			log.Fatalf("Error, marshalling JSON: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func getTheFact(db *sql.DB, id int) (database.Item, error) {
	if id == 0 {
		id = rand.Intn(539)
	}

	return database.GetItem(db, id), nil
}
