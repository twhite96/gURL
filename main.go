package main

import (
  "net/http"
  "log"
  "fmt"
  "io"
  "bufio"
  "os"
  "net/url"
)

var baseURL = "https://github.com";

var pl = fmt.Println;


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
  // Figure out how template strings work in go func() {
    
  pl("What repo do you want to GET?")

  reader := bufio.NewReader(os.Stdin)
  repo, err := reader.ReadString('\n')

  if err == nil {
    pl("Getting ", repo + "from GitHub")
  } else {
    log.Fatal(err)
  }

} 
