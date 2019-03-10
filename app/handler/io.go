package handler

import (
	"io"
	"flag"
	"mime"
	"fmt"
	"time"
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
	
	corsHeader := flag.String("allow-origin", "*", "value for Access-Control-Allow-Origin header")
	maxAge := flag.Int("max-age", 31557600, "Lifetime (in seconds) for "+
		"setting Cache-Control and Expires headers.  Defaults to one year.")
	
	flag.Parse()
	
	//
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
	
	//set corsHeader
	w.Header().Set("Access-Control-Allow-Origin", corsHeader)
	time.D
	//set Expiry Header
	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
	expiration := time.Now().Add(time.Duration(*maxAge) * time.Second)
	w.Header().Set("Expires", expiration.Format(time.RFC1123))
	
	ctype := getMimeType(file)
	w.Header().Set("Content-Type", ctype)
	
	w.Header().Set("Content-Length", strconv.FormatInt(file.Size(), 10))
	w.Header().Set("ETag", file.MD5())
	//done set Header
	
	io.Copy(w, file)
} 

func SaveFile(w http.ResponseWriter, r *http.Request) {
	//TODO
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
