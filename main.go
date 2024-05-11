package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//type "curl localhost:8080" and it will use your ip to show the weather in your area.
//type "curl localhost:8080/buenos_aires" and it will show the weather in that location.

func weatherHandler() {
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/{loc}", homepage)
	newRouter.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", newRouter))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["loc"]
	if location == "" {
		location = "auto:ip"
	}

	currentWeather, err := getWeather(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	printWeather(w, currentWeather)
}

func main() {
	weatherHandler()
}
