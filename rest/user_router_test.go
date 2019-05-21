package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type request struct {
	route string
	body  string
}

type Client struct {
	token string
}

var client Client

func (c *Client) login() (ok bool) {
	type response struct {
		Token string
	}
	body, err := c.makeRequest(request{
		`/users/login`,
		`{"Email": "olexiy.tkachenko+3@gmail.com","Passwoerd": "testpassword"}`,
	})
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
		c.token = res.Token
		return true
	}
	return false
}

func (c *Client) makeRequest(r request) ([]byte, error) {
	endpoint := "http://localhost:8000/api"
	url := endpoint + r.route

	payload := strings.NewReader(r.body)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

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

func TestLogin(t *testing.T) {
	if !client.login() {
		t.Fail()
	}
}
