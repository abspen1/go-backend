package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abspen1/restful-go/webapp"

	"github.com/gorilla/mux"
)

/*
*	Simple RESTful API created with GOlang
*	This is using localhost:8080
*
*
*
 */

func allProjects(w http.ResponseWriter, r *http.Request) {
	s := webapp.GetString()

	json.NewEncoder(w).Encode(s)
}

func postProjects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj webapp.Project
	_ = json.Unmarshal(info, &proj)
	webapp.SetString(proj)
	fmt.Fprintf(w, string(info))

	fmt.Fprintf(w, "Test POST endpoint worked!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/projects", allProjects).Methods("GET")
	myRouter.HandleFunc("/projects", postProjects).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", myRouter))
}

func main() {
	handleRequests()
}
