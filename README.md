# Go Heroku Template

Go template for Heroku apps

[Medium Blog](https://medium.com/p/deploy-gin-gonic-rest-api-free-in-heroku-code-ae167b7d0d53)

```
# go.mod & go.sum can be removed at this point
go mod init server

# packages
go get -u github.com/gin-gonic/gin

# Running
go mod vendor #update dependencies
go run server.go
# Listening at my-server.herokuapp.com/ping
# Would receive a json
```


## Stress the server

To test and check time requests you can use Apache Bench

```
sudo apt-get install apache2-utils

ab -t 10 -n 10000 -c 100 http://localhost:8080/ping
# or
ab -t 10 -n 10000 -c 100 http://localhost:8080/pingTime
```
