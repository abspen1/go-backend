package email

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"

	"github.com/badoux/checkmail"
)

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Info struct
type Info struct {
	Name    string
	Email   string
	Message string
}

// BotsFFL struct
type BotsFFL struct {
	Name  string
	Email string
	State string
}

// Birthday struct
type Birthday struct {
	Name          string
	Email         string
	JokeSetup     string
	JokePunchLine string
	Auth          string
}

// GetBdayEmail func that will display info on this endpoint
func GetBdayEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Austin's bday emailer post address :)")
}

// GetEmail func to display simple email endpoint
func GetEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Send email endpoint, nothing to see here!")
}

// PostEmail func to send email
func PostEmail(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info Info
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

	if SendEmail(info) {
		fmt.Fprintf(w, "Email sent successfully")
		return
	}
	fmt.Fprintf(w, "Email not sent")
}

// PostBdayEmail func for sending bday reminder email
func PostBdayEmail(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var info Birthday
	_ = json.Unmarshal(body, &info)
	resp := SendBdayEmail(info)
	if resp == "Success" {
		fmt.Fprintf(w, "Email sent successfully")
		return
	}

	if resp == "Auth Err" {
		fmt.Fprintf(w, "Invalid Authentification")
		return
	}

	fmt.Fprintf(w, "Email not sent")
}

// SendEmail function
func SendEmail(info Info) bool {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	// Receiver Email address.
	to := []string{
		from,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Name: %s  Email: %s Message: %s", info.Name, info.Email, info.Message)
	msg := "From: " + from + "\n" + "Subject: webapp\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

// SaveBotsInfo function
func SaveBotsInfo(info BotsFFL) bool {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	// Receiver Email address.
	to := []string{
		from,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Name: %s  Email: %s State: %s", info.Name, info.Email, info.State)
	msg := "From: " + from + "\n" + "Subject: BotsFFL\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Email Sent!")
	return true
}

// SendBdayEmail function
func SendBdayEmail(info Birthday) string {
	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL-PASS")
	pass := os.Getenv("BACK-END-AUTH")
	if info.Auth != pass {
		return "Auth Err"
	}
	// Receiver Email address.
	to := []string{
		from, info.Email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Happy birthday %s! Here is a joke to get your day started right!\n\n\n%s\n\n%s\n\n\n\n\nBest Regards,\n\nAustin", info.Name, info.JokeSetup, info.JokePunchLine)
	msg := "From: " + from + "\n" +
		"To: " + info.Email + "\n" +
		"Subject: Happy Birthday\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return "Email Err"
	}
	fmt.Println("Email Sent!")
	return "Success"
}
