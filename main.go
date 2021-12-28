package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/SarthakJain26/go-microservice/details"
	"github.com/SarthakJain26/go-microservice/geometry"

	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing health request")
	health := map[string]string{
		"STATUS": "UP",
		"TIME":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(health)
}

func UselessHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing useless info")
	area := geometry.AreaRectangle(rand.Float64(), rand.Float64())
	perimeter := geometry.Perimeter(rand.Float64(), rand.Float64())
	fmt.Fprintf(w, "Here is area and perimeter of random rectangles %f %f", area, perimeter)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing root request")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "At your service king %s ", r.URL.Query().Get("name"))
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing details request")
	hostname, _ := details.GetHostName()
	ip, _ := details.GetIP()
	details := map[string]string{
		"Hostname": hostname,
		"IP":       ip.String(),
	}
	json.NewEncoder(w).Encode(details)
}

func main() {
	fmt.Println("Hey Sarthak! All the best for your 1st microservice")
	// Instance of mux router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/details", DetailsHandler).Methods("GET")
	r.HandleFunc("/useless", UselessHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
