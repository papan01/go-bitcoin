package server

import (
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

//初始化router與設置routes，設置middlewares
func (t *App) initialize() {
	t.router = mux.NewRouter().StrictSlash(true)
	t.setupRoutes()
	//t.router.Use(limit)
}

//設置靜態網頁路徑，與註冊handlers
func (t *App) setupRoutes() {
	t.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	registerHandlers(t.router)
	t.router.Use(limit)
}

//啟動server
func (t *App) Run() {
	t.initialize()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, t.router))
}
