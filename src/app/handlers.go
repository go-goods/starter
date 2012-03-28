package main

import (
	"log"
	"net/http"
	"path"
	"strconv"
)

func handle_index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		perform_status(w, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "text/html")
	execute(w, tmpl_root("blocks", "index.block"))
}

//make a silly handler for testing statuses
func handle_status(w http.ResponseWriter, req *http.Request) {
	status, err := strconv.ParseInt(path.Base(req.URL.Path), 10, 32)
	if err != nil {
		log.Println(err)
		perform_status(w, http.StatusNotFound)
		return
	}
	perform_status(w, int(status))
}
