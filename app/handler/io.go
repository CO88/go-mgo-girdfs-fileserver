package handler

import (
	"io"
//	"flag"
	"mime"
	"fmt"
//	"time"
	"bufio"
	"errors"
	"strconv"
	"mime/multipart"
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
	for _, fileHeaders := range r.MultipartForm.File {
        for _, fileHeader := range fileHeaders {
            file, _ := fileHeader.Open()            
            if gridFile, err := db.Gridfs.Create(fileHeader.Filename); err != nil {
                //errorResponse(w, err, http.StatusInternalServerError)
                return
            } else {
                //gridFile.SetMeta(fileMetadata)
                gridFile.SetName(fileHeader.Filename)
                if err := writeToGridFile(file, gridFile); err != nil {
                    //errorResponse(w, err, http.StatusInternalServerError)
                    return
                }
            }
        }
	}
}

func writeToGridFile(file multipart.File, gridFile *mgo.GridFile) error {
	reader := bufio.NewReader(file)
	defer file.Close()
	
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New("Could not read the input file")
		}
		if n == 0 {
			break
		}
		//write a chunk
		if _, err := gridFile.Write(buf[:n]); err != nil {
			return errors.New("Could not write to GridFs for "+ gridFile.Name())
		}
	}
	gridFile.Close()
	return nil
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
