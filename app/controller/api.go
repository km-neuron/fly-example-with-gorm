package controller

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/rg-app/app/handler"
)

type API struct {
	studentHandle handler.StudentHandler
	embed         embed.FS
	mux           *http.ServeMux
}

func (api *API) BaseViewPath() *template.Template {
	var tmpl = template.Must(template.ParseFS(api.embed, "app/views/*"))
	return tmpl
}

func NewAPI(studentHandle handler.StudentHandler, embed embed.FS) API {
	mux := http.NewServeMux()
	api := API{
		studentHandle,
		embed,
		mux,
	}

	mux.HandleFunc("/", api.IndexPage)

	mux.Handle("/api/student/add", http.HandlerFunc(api.AddStudent))
	mux.Handle("/api/student/read", http.HandlerFunc(api.ReadStudent))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":"+port, api.Handler())
}
