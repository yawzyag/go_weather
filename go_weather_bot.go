package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yesid/go_weather_bot/handlers"
	"github.com/yesid/go_weather_bot/utils"
)

func main() {
	port := 8080
	http.HandleFunc("/", handlers.SayHello)
	http.HandleFunc("/healt", handlers.HealthCheckHandler)
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	log.Println("listen on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), utils.LogRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
