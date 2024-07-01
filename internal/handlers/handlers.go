package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/Kontentski/weatherly/internal/weather"
	"github.com/gorilla/mux"
)

func New() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/{loc}", homepage)
	newRouter.HandleFunc("/", homepage)
	newRouter.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("../../public/"))))

	log.Fatal(http.ListenAndServe(":8080", newRouter))
}

func homepage(w http.ResponseWriter, r *http.Request) {

	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = r.RemoteAddr
	}

	vars := mux.Vars(r)
	location := vars["loc"]
	if location == "" {
		location = clientIP
	}

	currentWeather, err := weather.Get(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		// Render HTML template

		tmpl := template.New("index.html").Funcs(template.FuncMap{
			"UnixTimeIsPast": weather.UnixTimeIsPast,
			"UnixTimeFormat": weather.UnixTimeFormat,
			"IntTemp":        weather.IntTemp,
		})
		tmpl, err := tmpl.ParseFiles("../../internal/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, currentWeather)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Render plain text
		w.Header().Set("Content-Type", "text/plain")
		weather.Print(w, currentWeather)
	}
}
