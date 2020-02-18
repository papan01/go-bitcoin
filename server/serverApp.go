package server

import (
	"bitcoin/modules"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//程式主結構:
//router: a request router
type App struct {
	router *mux.Router
}

//初始化router與設置routes
func (t *App) initialize() {
	t.router = mux.NewRouter().StrictSlash(true)
	t.setupRoutes()
}

//設置靜態網頁路徑，與將module中RoutesMap註冊handler
func (t *App) setupRoutes() {
	t.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	var module modules.Module
	module.Initialize()
	t.registerHandlers(module.RoutesMap)
}

//註冊handles
func (t *App) registerHandlers(src map[string]func(http.ResponseWriter, *http.Request)) {
	for key, value := range src {
		t.router.HandleFunc(key, value)
	}
}

//啟動server
func (t *App) Run() {
	t.initialize()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, limit(t.router)))
}
