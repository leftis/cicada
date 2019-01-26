package server

import (
"github.com/julienschmidt/httprouter"
"github.com/leftis/cicada/configuration"
"html/template"
"log"
"net/http"
)

var entryTemplate *template.Template

type HTML struct {}

func Init(app configuration.App) {
	var err error

	entryTemplate, err = template.ParseFiles(app.CurrentDirectory + "/server/admin-front.html")
	if err != nil {
		panic(err)
	}
	router := httprouter.New()
	router.GET("/admin", Admin)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	err := entryTemplate.Execute(w, HTML{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
