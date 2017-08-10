package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Ariticle 资产结构
type Ariticle struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Ariticles 资产集描述
type Ariticles []Ariticle

func returnAllAriticles(w http.ResponseWriter, r *http.Request) {
	ariticles := Ariticles{
		Ariticle{Title: "t1", Desc: "第一个", Content: "其实我也不知道是什么"},
		Ariticle{Title: "t2", Desc: "第二个", Content: "这是第二个什么鬼。。"},
	}
	fmt.Println("hit this all..")
	json.NewEncoder(w).Encode(ariticles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllAriticles)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my first http app")
	fmt.Println("response...")
}

func main() {
	handleRequests()
}
