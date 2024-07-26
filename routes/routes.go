package routes

import (
	"banner-art-web/ascii"
	"net/http"
)

func RoutesSetUp(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", ascii.IndexHandler)
	mux.HandleFunc("/ascii-art", ascii.SubmitHandler)
	mux.HandleFunc("/export", ascii.ExportHandler)
}
