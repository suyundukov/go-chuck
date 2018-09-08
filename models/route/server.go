package route

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// StartServer starts server
func StartServer(db *sql.DB) {
	r := httprouter.New()

	r.ServeFiles("/public/static/*filepath", http.Dir("public/static/"))
	r.GET("/favicon.ico", faviconHandler())
	r.GET("/fact/:id", idHandler(db))
	r.GET("/api/:id", apiHandler(db))
	r.GET("/api/", apiHandler(db))
	r.GET("/", mainHandler(db))

	log.Fatal(http.ListenAndServe(":"+initPort(), r))
}

func initPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}
