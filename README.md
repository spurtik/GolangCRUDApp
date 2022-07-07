# GolangCRUDApp
** Clone the repo to your local machine **

## Fork the Repository
Repo successfully forked at https://github.com/spurtik/GolangCRUDApp.git

## Clone down your fork of the repo locally. 
You can copy and paste either version of this line in to your terminal.
image.png


## Run the code

From the command line in the directory containing main.go, run the code to start the server.

$ go run main.go
From a different command line window, use curl to make a request to your running web service.
 
## Get all users

$ curl http://localhost:8080/users

## Create user 
$ curl http://localhost:8080/user \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "test1","email": "test1@xyz.com"}'


 ## Get user by ID

 $ curl http://localhost:8080/user/:id

 ## Update user 

 $ curl http://localhost:8080/user/:id \
--include \
    --header "Content-Type: application/json" \
    --request "PATCH" \
    --data '{"name": "test2","email": "test2@xyz.com"}'


## Delete user

$ curl -X DELETE http://localhost:8080/user/:id -H "Accept: application/json"

   
