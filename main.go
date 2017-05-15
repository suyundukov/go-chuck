// +build !appengine

package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type dataBase struct {
	Value []fact `json:"value"`
}

type fact struct {
	ID   float64 `json:"id"`
	Fact string  `json:"fact"`
}

var (
	t  = template.Must(template.ParseFiles("public/templates/index.tmpl"))
	db dataBase
)

func main() {
	r := httprouter.New()

	r.ServeFiles("/public/static/*filepath", http.Dir("public/static/"))
	r.GET("/favicon.ico", faviconHandler)
	r.GET("/fact/:id", idHandler)
	r.GET("/api", apiHandler)
	r.GET("/", mainHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	db = loadData()
}

func loadData() dataBase {
	data, err := ioutil.ReadFile("data/db.json")
	if err != nil {
		log.Fatalf("Error, loading database: %v", err)
	}

	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Fatalf("Error, parsing database: %v", err)
	}

	return db
}

func getTheFact(id int) (fact, error) {
	if id == 0 {
		id = rand.Intn(539)
	}

	return db.Value[id], nil
}

func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/static/favicon.ico")
}

func idHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id, err := strconv.Atoi(id); id < 540 && err == nil {
		fact, _ := getTheFact(id)
		t.ExecuteTemplate(w, "index.tmpl", fact)
	} else {
		http.Error(w, "404 page not found", 404)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fact, _ := getTheFact(0)
	t.ExecuteTemplate(w, "index.tmpl", fact)
}

func apiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fact, _ := getTheFact(0)
	res, err := json.Marshal(fact)
	if err != nil {
		log.Fatalf("Error, marshalling JSON: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
