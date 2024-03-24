package main

import (
  "net/http"
  "log"
  "context"
	"encoding/json"
  "fmt"
  "io"
  "net/url"
)

var baseURL = "https://github.com";

// Started to use a Client class and instead found this
// excellent answer on StackOverflow
// https://stackoverflow.com/a/71181938/3800146


// Google's Go client uses a POST req
// Here, we will authenticate and make a GET request



func (c *Client) Authenticate(username, password) error {
  setPassword := func(values url.Values) {
    values.Set("login", username)
    values.Set("password", password)
  }
}



func main() {
  // Figure out how template strings work in Go
}


