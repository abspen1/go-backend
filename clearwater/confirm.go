package clearwater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// PostConfirm func to send trip confirmation
func PostConfirm(w http.ResponseWriter, r *http.Request) {
	var info []byte

	if r.Body != nil {
		defer r.Body.Close()
		info, _ = ioutil.ReadAll(r.Body)
	}
	var trip Trip
	_ = json.Unmarshal(info, &trip)
	sendConfirmation(trip)
	fmt.Fprintf(w, "Success")
}

func sendConfirmation(trip Trip) {
	url := "https://api.pepipost.com/v5/mail/send"
	key := os.Getenv("api-key")

	// payload := strings.NewReader("{\"from\":{\"email\":\"clearwater.scheduling@austinspencer.works\",\"name\":\"Trip Confirmation\"},\"subject\":\"Your Clearwater trip info\",\"content\":[{\"type\":\"html\",\"value\":\"Hello [%NAME%], your trip is scheduled for dates [%START%] to [%END%].\"},{\"type\":\"html\",\"value\":\"Be on the lookout for an email with further information when the start date of your trip gets closer.\"}],\"personalizations\":[{\"attributes\": {\"NAME\": \"" + trip.Name + "\",\"START\": \"" + trip.StartDate + "\",\"END\": \"" + trip.EndDate + "\"}\"to\":[{\"email\":\"" + trip.Email + "\",\"name\":\"" + trip.Name + "\"}]}]}")
	payload := strings.NewReader("{\"from\":{\"email\":\"clearwater.scheduling@austinspencer.works\",\"name\":\"Trip Confirmation\"},\"subject\":\"Your Clearwater trip info\",\"content\":[{\"type\":\"html\",\"value\":\"Hello [%NAME%], your trip is scheduled for dates [%START%] to [%END%].\"}],\"personalizations\":[{\"attributes\":{\"NAME\":\"" + trip.Name + "\",\"START\":\"start date\",\"END\":\"end date\"},\"to\":[{\"email\":\"abspencer2097@gmail.com\",\"name\":\"Austin Spencer\"}]}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("api_key", key)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Print("Res: ")
	fmt.Println(res)
	fmt.Print("Body: ")
	fmt.Println(string(body))
}

