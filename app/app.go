package app

import (
	"net/http"
	"fmt"
	"log"
	
	"github.com/gorilla/mux"
	"github.com/blanccobb/go-mgo-girdfs-fileserver/app/handler"
	"github.com/blanccobb/go-mgo-girdfs-fileserver/app/db"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	
	db.Init()
	
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/", a.GetRoot)
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

// Todo Handler
func (a *App) GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
	handler.LoadUploadPage(w, r)
}


func (a *App) DownloadFile(w http.ResponseWriter, r *http.Request) {
	handler.GetFile(w,r)
}

func (a *App) UploadFile(w http.ResponseWriter, r *http.Request) {
	handler.SaveFile(w,r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}