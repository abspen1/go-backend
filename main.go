package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abspen1/restful-go/webapp"

	"github.com/gorilla/handlers"
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
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handlers.CORS(originsOk, headersOk, methodsOk)(myRouter)))
}

func main() {
	handleRequests()
}
