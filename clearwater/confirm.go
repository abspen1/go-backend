package clearwater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	res := sendConfirmation(trip)
	fmt.Fprintf(w, res)
}

func sendConfirmation(trip Trip) string {
	url := "https://api.pepipost.com/v5/mail/send"
	key := os.Getenv("api-key")

	payload := strings.NewReader("{\"from\":{\"email\":\"clearwater.scheduling@austinspencer.works\",\"name\":\"Trip Confirmation\"},\"subject\":\"Your Clearwater trip info\",\"template_id\":28952,\"content\":[{\"type\":\"html\",\"value\":\"[%NAME%]\"}],\"personalizations\":[{\"attributes\":{\"NAME\":\"" + trip.Name + "\",\"START\":\"" + trip.StartDate + "\",\"END\":\"" + trip.EndDate + "\",\"ID\":\"" + trip.Password + "\"},\"to\":[{\"email\":\"" + trip.Email + "\",\"name\":\"" + trip.Name + "\"}]}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("api_key", key)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error"
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return "Error"
	}
	return "Success"
}

