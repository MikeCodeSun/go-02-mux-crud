package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Director *Director
}

type Director struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var moives []Movie



func getAllMoives(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(moives)
	
}

func helloHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}

func  getOneMoive(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err !=nil {
		log.Fatal(err)
	}
	
	for _, item := range moives {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func deleteMoive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "append/json")
	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	for index, item := range moives {
		if item.ID == id {
			moives = append(moives[:index], moives[index+1:]...)

		}
	}
	json.NewEncoder(w).Encode(moives)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var moive Movie
	if err := json.NewDecoder(r.Body).Decode(&moive); err != nil {
		fmt.Println(err)
	}
	moives = append(moives, moive)
	json.NewEncoder(w).Encode(moives)

}

// update moive
func updateFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var moive Movie
	json.NewDecoder(r.Body).Decode(&moive)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil{
		log.Fatal(err)
	}
	for index, item := range moives{
		if item.ID == id {
			moives = append(moives[:index], moives[index+1:]... )
		}
	}
	moives = append(moives, moive)
	json.NewEncoder(w).Encode(moives)
}

func main() {
	moives = append(moives, Movie{ID: 1, Name: "bad boy", Director: &Director{Name: "lee", ID: 1}})
	moives = append(moives, Movie{ID: 2, Name: "good girl", Director: &Director{Name: "Jack", ID: 2}})

	r := mux.NewRouter()
	r.HandleFunc("/", helloHome).Methods("GET")
	r.HandleFunc("/moive", getAllMoives).Methods("GET")
	r.HandleFunc("/moive/{id}", deleteMoive).Methods("DELETE")
	r.HandleFunc("/moive/{id}", getOneMoive).Methods("GET")
	r.HandleFunc("/moive", createMovie).Methods("POST")
	r.HandleFunc("/moive/{id}", updateFunc).Methods("PATCH")
	fmt.Printf("the movie is %v", moives)
	fmt.Println("go run server")
	log.Fatal(http.ListenAndServe(":8000", r))
}