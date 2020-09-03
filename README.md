# Go RESTful API
Building a RESTful API using Go programming language. I love the combination of Go and Redis as my backend for my website. I think Go is an awesome programming language and combined with the ease and speed of Redis is hard to beat. I would love to implement something more complex in the backend that would allow me to use Go Routines and channeling to get the most out of Go's speed and possibilities. This program is used in my [website](https://abspen1.github.io) as the backend for a few different pages. [Webapp repo](https://github.com/abspen1/abspen1.github.io)!

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

## 📁 email
* Package for the contact page on my website
* Imports the net/smtp package to allow me to send gmail of the contact info
* Will send an email in format:
```bash
webapp


Name: Joe  Email: example@host.com
Message: This is an example.
```

## 📁 test
* Messy code that I use to test specific functions, mostly with Redis and JSON

## Checkmail Package
* Checking to make sure the email given is a valid email
* Already a simple format check on the front end to catch obvious errors
    * This is a much more ellaborate check
    
## Running
### Run main.go
* cd to your go-backend directory
* In terminal: $ go run main.go
* Now your program/server is running

### Test GET/POST requests
* Using an app of your choosing (I use [Postman](https://www.postman.com/downloads/))
* Put in your request URL (localhost:8080/projects)
* Either choose POST or GET request and check the body of the output

## Built to run with Docker
### Docker Commands
* cd into the working directory

```bash
docker build -t imageName

docker run -d \
--name containerName \
--restart unless-stopped \
-p 8080:8080 \
imageName
```
    
## Docker Help
* $ docker --help
* $ docker image --help
* $ docker container --help
