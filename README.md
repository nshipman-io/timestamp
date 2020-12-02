# Timestamp 
Timestamp is a microservice built using Go. Based on the FreeCodeCamp JS Project: https://www.freecodecamp.org/learn/apis-and-microservices/apis-and-microservices-projects/timestamp-microservice

## Installation
1. go build

## Usage 
1. ./timestamp
2. From your browser visit the url 0.0.0.0:8080 or 127.0.0.1:8080
3. Enter a time string into the path /api/timestamp/{string}.
	a. http://0.0.0.0:8080/api/timestamp/1606873242
	b. http://0.0.0.0:8080/api/timestamp/2015-12-25
4. Enter no string in /api/timestamp/ to receive the current time.
