package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abspen1/restful-go/players"

	"github.com/abspen1/restful-go/alp"
	"github.com/abspen1/restful-go/home"
	"github.com/abspen1/restful-go/twitter/tweet"

	"github.com/abspen1/restful-go/players/rosters"
	"github.com/abspen1/restful-go/players/trending"

	"github.com/abspen1/restful-go/todos"

	"github.com/abspen1/restful-go/botsffl"

	"github.com/abspen1/restful-go/twitter"

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

	rpsUser = rps.SaveData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

func postRPSLogin(w http.ResponseWriter, r *http.Request) {
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

	rpsUser = rps.GetData(rpsUser)

	json.NewEncoder(w).Encode(rpsUser)
}

func getRPSLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game login endpoint, nothing to see here!")
}

func getRPS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rock Paper Scissors game save endpoint, nothing to see here!")
}

func getEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Send email endpoint, nothing to see here!")
}

func getRmprojects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove projects endpoint, nothing to see here!")
}

func getTwitterData(w http.ResponseWriter, r *http.Request) {
	data := twitter.GetTwitterData()

	json.NewEncoder(w).Encode(data)
}

func postTweet(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var content tweet.Content
	_ = json.Unmarshal(body, &content)
	if content.Auth != os.Getenv("SECRET") {
		fmt.Fprintf(w, "Invalid Authentification")
		return
	}
	resp := tweet.Tweet(content)
	if resp == true {
		fmt.Fprintf(w, "Tweet sent successfully")
	} else {
		fmt.Fprintf(w, "Error in postTweet")
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := todos.GetTodos()

	json.NewEncoder(w).Encode(todos)
}

func postTodos(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info todos.Todos
	_ = json.Unmarshal(body, &info)
	if todos.AddTodo(info) {
		fmt.Fprintf(w, "Added todo successfully")
	}

	fmt.Fprintf(w, "Todo not added due to an error")
}
func rmTodos(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info todos.FullTodo
	_ = json.Unmarshal(body, &info)
	if todos.RmTodo(info) {
		fmt.Fprintf(w, "Removed todo successfully")
	}

	fmt.Fprintf(w, "Todo wasn't removed due to an error")
}

func handleRequests() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://abspen1.github.io", "https://austinspencer.works"},
		AllowCredentials: true,
		Debug:            false,
	})
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/austinapi/", home.Page).Methods("GET")
	myRouter.HandleFunc("/austinapi/bdayemail", email.PostBdayEmail).Methods("POST")
	myRouter.HandleFunc("/austinapi/bdayemail", email.GetBdayEmail)
	myRouter.HandleFunc("/austinapi/botsffl", botsffl.GetBotsFFL).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl", botsffl.PostBotsFFL).Methods("POST")
	myRouter.HandleFunc("/austinapi/botsffl/teams/midwest", players.GetMwTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/west", players.GetWTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/northeast", players.GetNeTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/southeast", players.GetSeTeams).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/midwest/roster", rosters.GetMwRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/west/roster", rosters.GetWRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/northeast/roster", rosters.GetNeRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/teams/southeast/roster", rosters.GetSeRosters).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/daily/add", trending.GetDailyTrendAdd).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/daily/drop", trending.GetDailyTrendDrop).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/weekly/add", trending.GetWeeklyTrendAdd).Methods("GET")
	myRouter.HandleFunc("/austinapi/botsffl/trending/weekly/drop", trending.GetWeeklyTrendDrop).Methods("GET")
	myRouter.HandleFunc("/austinapi/current-stock-price", alp.Get).Methods("GET")
	myRouter.HandleFunc("/austinapi/current-stock-price", alp.PostStockPrice).Methods("POST")
	myRouter.HandleFunc("/austinapi/email", sendEmail).Methods("POST")
	myRouter.HandleFunc("/austinapi/email", getEmail).Methods("GET")
	myRouter.HandleFunc("/austinapi/go-tweet", postTweet).Methods("POST")
	myRouter.HandleFunc("/austinapi/go-tweet", tweet.GetTweet).Methods("GET")
	myRouter.HandleFunc("/austinapi/projects", getProjects).Methods("GET")
	myRouter.HandleFunc("/austinapi/projects", postProjects).Methods("POST")
	myRouter.HandleFunc("/austinapi/rmprojects", getRmprojects).Methods("GET")
	myRouter.HandleFunc("/austinapi/rmprojects", postRmprojects).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps/login", postRPSLogin).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps", postRPS).Methods("POST")
	myRouter.HandleFunc("/austinapi/rps/login", getRPSLogin).Methods("GET")
	myRouter.HandleFunc("/austinapi/rps", getRPS).Methods("GET")
	myRouter.HandleFunc("/austinapi/tendie-intern", getTwitterData).Methods("GET")
	myRouter.HandleFunc("/austinapi/todos", getTodos).Methods("GET")
	myRouter.HandleFunc("/austinapi/todos", postTodos).Methods("POST")
	myRouter.HandleFunc("/austinapi/todos/rm", rmTodos).Methods("POST")

	handler := c.Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8558", handler))
}

func main() {
	handleRequests()
}
