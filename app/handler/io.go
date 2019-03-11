package handler

import (
	"io"
//	"flag"
	"mime"
	"fmt"
//	"time"
	"strconv"
	"path/filepath"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/blanccobb/go-mgo-girdfs-fileserver/app/db"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	
	file, err := db.Gridfs.Open(name)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s Not Found\n", name)
			return 
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error\n")
		return 
	}
	defer file.Close()
	
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	
	ctype := getMimeType(file)
	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Content-Length", strconv.FormatInt(file.Size(), 10))
	w.Header().Set("ETag", file.MD5())
	
	io.Copy(w, file)
} 

func SaveFile(w http.ResponseWriter, r *http.Request) {
	//TODO
	//vars := mux.Vars(r)
	//name := vars["name"]
	
	
	
}

func getMimeType(file *mgo.GridFile) string {
	ctype := file.ContentType()
	
	if ctype == "" {
		ctype := mime.TypeByExtension(filepath.Ext(file.Name()))
		if ctype == "" {
			return "application/octet-stream"
		}
		return ctype
	}
	return ctype
}
