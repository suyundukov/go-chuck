package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DataBase is the actual db.json
type DataBase struct {
	Value []Fact `json:"value"`
}

// Fact is the actual Chuck Norris fact
type Fact struct {
	ID   float64 `json:"id"`
	Fact string  `json:"fact"`
}

var (
	templates = template.Must(template.ParseFiles("templates/index.tmpl"))
	db        DataBase
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

func loadData() DataBase {
	data, err := ioutil.ReadFile("data/db.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = db.UnmarshalJSON(data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return db
}

func getTheFact(id int) (Fact, error) {
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
	if id, err := strconv.Atoi(id); id < 539 && err == nil {
		fact, _ := getTheFact(id)
		templates.ExecuteTemplate(w, "index.tmpl", fact)
	} else {
		http.Error(w, "404 page not found", 404)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fact, _ := getTheFact(0)
	templates.ExecuteTemplate(w, "index.tmpl", fact)
}

func apiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fact, _ := getTheFact(0)
	res, err := fact.MarshalJSON()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
