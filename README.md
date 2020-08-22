# Go RESTful API
Building a RESTful API using Go programming language. This has been transitioned to be used in my website! This is the backend for my 'Next Project' page.

## Running
### Run main.go
* cd to your go-backend directory
* In terminal: $ go run main.go
* Now your program is running

### Test GET/POST requests
* Using an app of your choosing (I use [Postman](https://www.postman.com/downloads/))
* Put in your request URL (localhost:8081/articles)
* Either choose POST or GET request and check the body of the output

## Built to run with Docker
### Docker Commands
* cd into your directory
* $ docker build -t imageName
* $ docker run --name containerName -d -p 8080:8080 imageName

## 📁 webapp
### This is an imported folder I use to connect to redis
* The main bulk of working with get/post requests