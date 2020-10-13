package tweet

import (
	// other imports
	"fmt"
	"log"

	"github.com/abspen1/restful-go/auth"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Content struct is the content that we expect in the post request
type Content struct {
	Message string
	Auth    string
}

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

// Tweet function will tweet out what is going on
func Tweet(content Content) bool {
	// fmt.Println("Go-Twitter Bot v0.01")
	creds := Credentials{
		AccessToken:       auth.GoDotEnvVariable("KEY"),
		AccessTokenSecret: auth.GoDotEnvVariable("SECRET"),
		ConsumerKey:       auth.GoDotEnvVariable("CONSUMER_KEY"),
		ConsumerSecret:    auth.GoDotEnvVariable("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
		return false
	}

	message := content.Message + "\n#GoBot"

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	_, _, err = client.Statuses.Update(message, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}