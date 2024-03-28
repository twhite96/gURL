package user

import(
  "net/url"
)


func (c *Client) Authenticate(username, password) error {
  setPassword := func(values url.Values) {
    values.Set("login", username)
    values.Set("passwor, password)
  }
}
