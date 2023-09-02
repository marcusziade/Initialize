package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *http.ServeMux
}

func New() *App {
	return &App{
		Router: http.NewServeMux(),
	}
}

func (a *App) Run() {
	a.initializeRoutes()
	a.initializeDatabase()
	a.initializeServer()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/users", a.handleUsers)
}

func (a *App) initializeDatabase() {
	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		file, err := os.Create("database.db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}

func (a *App) initializeServer() {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         "",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}

func (a *App) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.getUsers(w, r)
	case http.MethodPost:
		a.createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create user")
}
