package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Ariticle 资产结构
type Ariticle struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Ariticles 资产集描述
type Ariticles []Ariticle

var ariticles = Ariticles{
	Ariticle{Title: "t1", Desc: "第一个", Content: "其实我也不知道是什么"},
	Ariticle{Title: "t2", Desc: "第二个", Content: "这是第二个什么鬼。。"},
}

func returnAllAriticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hit this all..")
	json.NewEncoder(w).Encode(ariticles)
}

func singleAriticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	fmt.Fprintf(w, "key: %d => %+v", key, ariticles[key])
}

func postAriticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appalication-json")

	var a Ariticle

	b, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b, &a); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	ariticles = append(ariticles, a)

	fmt.Fprintf(w, "OK: %+v", a)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllAriticles).Methods("GET")
	router.HandleFunc("/ariticle/{id}", singleAriticle).Methods("GET")
	router.HandleFunc("/ariticle", postAriticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my first http app")
	fmt.Println("response...")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
