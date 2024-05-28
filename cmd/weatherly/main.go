package main

import (
	"github.com/Kontentski/weatherly/internal/handlers"
)

//type "curl localhost:8080" and it will use your ip to show the weather in your area.
//actually if you run localhost it won't use your ip because it is one machine, so the output will be incorrect, use tools like ngrok to catch your ip.
//type "curl localhost:8080/buenos_aires" and it will show the weather in that location.


func main() {
	handlers.New()
}
