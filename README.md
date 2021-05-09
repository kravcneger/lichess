## Golang Lichess Package [![Go Report Card](https://goreportcard.com/badge/github.com/github.com/lukemilby/lichess)](https://goreportcard.com/report/github.com/github.com/lukemilby/lichess)

A Go implementation of [Lichess](https://lichess.org)'s [API](https://lichess.org/api)

### Usage

```go
package main

import (
 "fmt"
 "github.com/lukemilby/lichess"
 "log"
 "net/http"
 "net/url"
)

func main() {
 baseURL, err := url.Parse("https://lichess.org")
 if err != nil {
  log.Fatal(err)
 }
 client := new(lichess.Client)

 client.BaseURL = baseURL
 client.APIKey = "<API Key>"
 client.UserAgent = "Golang Client"
 client.HttpClient = new(http.Client)

 user, err := client.GetProfile()
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println(user.ID)
}
```

### Available Features

- [x] Accounts
- [ ] Users
- [x] Unit Testing
- [ ] Games
- [ ] Board
- [ ] Teams
- [ ] Tournaments
- [ ] Relations
- [ ] Challenges
- [ ] Chessbot

## 💻 Local Development

## Getting Started

- First, Create an API Token from [here](https://lichess.org/account/oauth/token/create)
- Second, [fork this repo](https://github.com/lukemilby/lichess/fork),
- Third, run these commands to clone it locally and get started

```zsh
# Clone and CD into/Open this project
$ git clone git@github.com:YOUR_GITHUB_USERNAME/lichess.git && cd lichess
# Create a .env file to hold our secrets
$ touch .env > LICHESS_TOKEN=YOUR_API_TOKEN
# Download & Install the dependancies. Then Compile the program
$ go build
# Run the program locally
$ go run lichess.go
```

## 📚 Resources

- [Go Template CheeatSheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)
- [Intro to go testing](https://tutorialedge.net/golang/intro-testing-in-go/)
- [Intro to go benchmark testing](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [Go: tests with HTML coverage report](https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d)
- [Makefiles for Go Developers](https://tutorialedge.net/golang/makefiles-for-go-developers/)
- [Runes in golang](https://www.geeksforgeeks.org/rune-in-golang/)
- [Default Vaules for Struct Fields in golang](https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/)
