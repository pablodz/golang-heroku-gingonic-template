# Go Heroku Template

Go template for Heroku apps

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
