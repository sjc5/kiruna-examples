package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sjc5/kiruna"
	global "github.com/sjc5/kiruna-examples/fullstack-minimal"
)

func main() {
	// Health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Serve static files from "dist/kiruna/static/public" directory, accessible at "/public/"
	http.Handle("/public/", global.Kiruna.GetServeStaticHandler("/public/"))

	// Serve an HTML file using html/template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		FS, err := global.Kiruna.GetPrivateFS()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFS(FS, "templates/index.go.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, struct {
			Kiruna     *kiruna.Kiruna
			FaviconURL string
		}{
			Kiruna:     global.Kiruna,
			FaviconURL: global.Kiruna.GetPublicURL("favicon.ico"),
		})
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
