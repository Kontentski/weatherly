package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//type "curl localhost:8080" and it will use your ip to show the weather in your area.
//actually if you run localhost it won't use your ip because it is one machine, so the output will be incorrect, use tools like ngrok to catch your ip.
//type "curl localhost:8080/buenos_aires" and it will show the weather in that location.

func weatherHandler() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/{loc}", homepage)
	newRouter.HandleFunc("/", homepage)
	newRouter.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

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

	currentWeather, err := getWeather(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		// Render HTML template

		tmpl := template.New("index.html").Funcs(template.FuncMap{
			"unixTimeFormat": unixTimeFormat,
			"unixTimeIsPast": unixTimeIsPast,
			"intTemp":        intTemp,
		})

		tmpl, err := tmpl.ParseFiles("/templates/index.html")
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
		printWeather(w, currentWeather)
	}
}

func main() {
	weatherHandler()
}
