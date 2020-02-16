package server

import (
	"bitcoin/modules"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func (t *App) Initialize() {
	t.router = mux.NewRouter().StrictSlash(true)
	var module modules.Module
	module.Initialize()
	t.addHandlers(module.APIMap)

	t.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	t.router.HandleFunc("/", Index)
}

func (t *App) addHandlers(src map[string]func(http.ResponseWriter, *http.Request)) {
	for key, value := range src {
		t.router.HandleFunc("/api"+key, value)
	}
}

func (t *App) Run() {
	t.Initialize()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, t.router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./static/index.html")
}
