# Go

[![N|Golang](https://miro.medium.com/fit/c/256/256/1*yh90bW8jL4f8pOTZTvbzqw.png)](https://golang.org/)

Go (also known as Golang) is an open source programming language developed by Google. It is a statically-typed compiled language. Go supports concurrent programming, i.e. it allows running multiple processes simultaneously. This is achieved using channels, goroutines, etc. Go has garbage collection which itself does the memory management and allows the deferred execution of functions.

## Weater bot

This application is made in order to show the current state of air quality and other states of climate in a specific city

### How to run

Make sure you have:

- the project cloned
- go installed (run "go version" on your terminal)

For runing the server, use this command file:

- go run go_weather_bot.go

For runining test:

- go test ./handlers
