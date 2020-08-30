package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abspen1/restful-go/email"

	"github.com/abspen1/restful-go/projects"
	"github.com/abspen1/restful-go/rps"
	"github.com/badoux/checkmail"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getProjects(w http.ResponseWriter, r *http.Request) {
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

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

func postRmprojects(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var proj projects.RmProject
	_ = json.Unmarshal(info, &proj)

	if projects.CheckPass(proj) {
		if projects.RmString(proj) {
			fmt.Fprintf(w, "POST remove worked!")
		} else {
			fmt.Fprintf(w, "Error")
		}
	} else {
		fmt.Fprintf(w, "Err")
	}

	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info email.Info
	_ = json.Unmarshal(body, &info)

	err := checkmail.ValidateFormat(info.Email)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}
	err = checkmail.ValidateHost(info.Email)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	if email.SendEmail(info) {
		fmt.Fprintf(w, "Email sent successfully")
		return
	}
	fmt.Fprintf(w, "Email not sent")
}

func postRPS(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var rpsUser rps.User
	_ = json.Unmarshal(info, &rpsUser)

	err := checkmail.ValidateFormat(rpsUser.Username)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Format Error")
		return
	}
	err = checkmail.ValidateHost(rpsUser.Username)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		fmt.Fprintf(w, "Error")
		return
	}

	rpsUser = rps.SaveData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

func getRPS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors backend")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Austin's API, nothing to see here!")
}

func handleRequests() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://abspen1.github.io", "https://abspen1.github.io/next-project/np.html"},
		AllowCredentials: true,
		Debug:            false,
	})
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/austinapi/", homePage)
	myRouter.HandleFunc("/austinapi/projects", getProjects).Methods("GET")
	myRouter.HandleFunc("/austinapi/projects", postProjects).Methods("POST")
	myRouter.HandleFunc("/austinapi/rmprojects", postRmprojects).Methods("POST")
	myRouter.HandleFunc("/austinapi/email", sendEmail).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps", getRPS).Methods("GET")
	myRouter.HandleFunc("/austinapi/rps", postRPS).Methods("POST")
	handler := c.Handler(myRouter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}

func main() {
	handleRequests()
}
