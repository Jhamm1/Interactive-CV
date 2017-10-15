package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var portnumer string = ":8082"
var results []string
var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", portnumer, "Port to listen on")
)

// Article - Our struct for all articles
type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//Equipment recommendation struct
type Equipment struct {
	PatientId     int    `json:"PatientId"`
	FurnitureItem string `json:"FurnitureItem"`
	EquipmentType string `json:"EquipmentType"`
	EquipmentSize string `json:"EquipmentSize"`
}

type Equipments []Equipment

func homescreen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//definition for endpoint
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

//definition for endpoint
func returnER(w http.ResponseWriter, r *http.Request) {
	equipments := Equipments{
		Equipment{PatientId: 11, FurnitureItem: "Toilet",
			EquipmentType: "Toilet raiser", EquipmentSize: "4-inch"},
		//Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(equipments)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func returnMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnMessages")
	json.NewEncoder(w).Encode(results)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/PID/equipment", returnER)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	//results = append(results, time.Now().Format(time.RFC3339))
	myRouter.HandleFunc("/post/1", returnMessages)
	myRouter.HandleFunc("/post", PostHandler)
	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(portnumer, myRouter))

}

func main() {
	handleRequests()
}
