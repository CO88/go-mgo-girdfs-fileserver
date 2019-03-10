package app

import (
	"net/http"
	"log"
	
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
}

func (a *App) setRouters() {
	a.Get("/fs/{name}", a.DownloadFile)
	a.Post("/fs/{name}", a.UploadFile)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) DownloadFile(w http.ResponseWriter, r *http.Request) {
	Handler.GetFile(w,r)
}

func (a *App) UploadFile(w http.ResponseWriter, r *http.Request) {
	Handler.SaveFile(w,r)
}
