package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Route string
	Body  string
}

type Client struct {
	Token string
}

type Credentials struct {
	Email    string
	Password string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Login(creds Credentials) (ok bool) {
	type response struct {
		Token string
	}
	js, _ := json.Marshal(creds)
	body, err := c.MakeRequest(Request{`/account/login`, string(js)})
	if err != nil {
		log.Fatal(err)
		return false
	}
	res := &response{}
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(res.Token) > 0 {
		c.Token = res.Token
		return true
	}
	return false
}

func (c *Client) EnsureLogin(creds Credentials) {
	if len(c.Token) == 0 {
		c.Login(creds)
	}
}

func (c *Client) Logout() {
	c.Token = ""
}

func (c *Client) Me() (me string) {
	me = ""
	body, err := c.MakeRequest(Request{Route: `/account/me`})
	if err != nil {
		log.Fatal(err)
		return
	}
	return string(body)
}

func (c *Client) MakeRequest(r Request) ([]byte, error) {
	endpoint := "http://localhost:8000/api"
	url := endpoint + r.Route

	payload := strings.NewReader(r.Body)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+c.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("non 2xx status: %v", res.Status))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
