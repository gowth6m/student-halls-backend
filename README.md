# Student Halls Backend

## Description

This is a backend for the Student Halls project. It is a RESTful API that allows users to create, read, update, and delete data from a MongoDB. The architecture is a serverless appraoch using Vercel's serverless functions. The API is built with Go using Go Gin Gonic and the MongoDB driver.

## How to run

1. Clone the repository
2. Run `go mod download` to download the dependencies
3. Run `go run cmd/server/main.go` or `make server` to start the server
4. The server will be running on `localhost:8080`

## Endpoints

View the API documentation [here](https://api.student-halls.com/swagger/index.html)