# Go Heroku Template

Go template for Heroku apps

[Medium Blog](https://medium.com/p/deploy-gin-gonic-rest-api-free-in-heroku-code-ae167b7d0d53)

```bash
# Download latest packages
go mod tidy
# Try on localhost
go run main.go
```


## Stress the server

To test and check time requests you can use Apache Bench

```bash
# Debian based package
sudo apt-get install apache2-utils

ab -t 10 -n 10000 -c 100 http://localhost:8080/ping
# or
ab -t 10 -n 10000 -c 100 http://localhost:8080/pingTime
```
