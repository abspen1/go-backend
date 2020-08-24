package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abspen1/restful-go/projects"
	"github.com/gorilla/mux"
)

func allProjects(w http.ResponseWriter, r *http.Request) {
	s := projects.GetString()

	json.NewEncoder(w).Encode(s)
}

func postProjects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj projects.Project
	_ = json.Unmarshal(info, &proj)
	projects.SetString(proj)
	fmt.Fprintf(w, string(info))

	fmt.Fprintf(w, "Test POST endpoint worked!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/austinapi/", homePage)
	myRouter.HandleFunc("/austinapi/projects", allProjects).Methods("GET")
	myRouter.HandleFunc("/austinapi/projects", postProjects).Methods("POST")
	handler := c.Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8558", handler))
}

func main() {
	handleRequests()
}
