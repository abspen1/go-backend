# Go RESTful API [![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fabspen1%2Fgo-backend&count_bg=%2329BEB0&title_bg=%23555555&icon=go.svg&icon_color=%2329BEB0&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)
Building a RESTful API using Go programming language. I love the combination of Go and Redis as my backend for my website. I think Go is an awesome programming language and combined with the ease and speed of Redis is hard to beat. I would love to implement something more complex in the backend that would allow me to use Go Routines and channeling to get the most out of Go's speed and possibilities. This program is used in my [website](https://abspen1.github.io) as the backend for a few different pages. Also check out my websites [repository](https://github.com/abspen1/abspen1.github.io)!


## 📁 auth (Deprecated)
* Package that imports goDotEnv
* This allows me to not need the goDotEnv function in each project and rather just import the function from my auth project


## 📁 botsffl
* Package that is pretty ugly to be honest
* Need to clean up this code but it works for now
* Back-end for my fantasy football league's hub [BOTS-FFL](https://austinspencer.works/BOTS-FFL)
* This package is how I get the standings and point leaders on the website


## 📁 email
* Package for the contact page on my website
* Imports the net/smtp package to allow me to send gmail of the contact info
* Will send an email in format:
```bash
webapp


Name: Joe  Email: example@host.com
Message: This is an example.
```
* **Now also sending birthday messages from Python script api request**


## 📁 players
* Populates team names for the dropdown bar within the Roster section of my website [here](https://austinspencer.works/BOTS-FFL/#rosters)
* Most of the work here is done on my Python scripts that pulls the data from Sleeper.app and saves it into my Redis database

### 📁 rosters
* Largest set of data on my back-end server so far
* Shows the full roster for all 12 teams within each region
* In total around 300 players for each region

### 📁 trending
* Package that shows the trending players based on adds/dropped in sleeper leagues
* Pretty cool and simple call
* The request is made on a Python script that then saves the top 10 for 24hour and 5 day ranges respectively
* All visible within the trending section of the website [here](https://austinspencer.works/BOTS-FFL/#trending-players)


## 📁 projects
* Package I use to connect to redis for 'Next Project' page in my website
* The main bulk of working with saving the project information
* The way I use Redigo in this package is actually pretty awesome..
   * Using a set I have the key as projects and the value added is a JSON Marshall of the struct Project : Project Description


## 📁 rps
* Rock Paper Scissors backend package
* Will save a user's score with a hash like this:
```bash
username {
      wins: 20,
      losses: 20,
}      
```

## 📁 test
* Messy code that I use to test specific functions, mostly with Redis and JSON


## 📁 todos
* This package displays my todos json list
* Made to be used on my website but haven't gotten around to actually implementing it
* Would be something to look into finishing at a later point


## 📁 twitter
* This is how I am displaying the data for my twitter bot on the webpage
* This is one of my all-time favorite projects I have ever done and I have spent a lot of time doing it
* Please check out the [webpage](https://austinspencer.works/twitter-bot) and the [source code](https://github.com/abspen1/twitter-bot)

### 📁 tweet
* This package uses Go to with the twitter API to tweet out specific tweets
* As of right now the only tweets being sent out with the back-end post request is about Call of Duty
* The main chunk of code is built with JavaScript and uses the call-of-duty-api npm package
* The JavaScript application will tweet out every hour if there are new stats to log of the specific users (myself and 1 other friend currently)
* Will also tweet out the weekly stats at the end of the week
* I'm looking forward to taking advantage of this package more in the future!!


## Checkmail Package
* Checking to make sure the email given is a valid email
* Already a simple format check on the front end to catch obvious errors
    * This is a much more ellaborate check
    

# Running

## Run main.go
* cd to your go-backend directory
* In terminal: $ go run main.go
* Now your program/server is running

## Test GET/POST requests
* Using an app of your choosing (I use [Postman](https://www.postman.com/downloads/))
* Put in your request URL (localhost:8558/projects)
* Either choose POST or GET request and check the body of the output

## Built to run as Docker container
### Docker Commands
* cd into the working directory


```bash
# Build the image
docker build -t imageName

# Run the image you just built as a container
docker run -d \
--name containerName \
--restart unless-stopped \
# Expose the port
-p 8558:8558 \
imageName
```

## Kubernetes Implementation
* Added austin-api-dc-svc-ingress.yaml which is for the Kubernetes deployment
* Integrated from the .env into env variables saved in Kubernetes secrets
* Something to keep in mind is never use .env files with Docker containers since then anyone running the docker container will have access to those .env variables

### Rancher Pipeline
* Within Kubernetes, have pipelining set up with this GitHub repository
    * Each push in the repository will redeploy the project with the changes implemented
    
## Docker Help
* $ docker --help
* $ docker image --help
* $ docker container --help
