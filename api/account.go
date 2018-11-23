package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Account struct {
	client  *http.Client
	cookies []*http.Cookie
}

type result struct {
	Code    int    `json:"code"`
	Items   []Item `json:"items"`
	Message string `json:"message"`
}

func NewAccount() *Account {
	return &Account{client: &http.Client{}}
}

func (a *Account) Login(username, password string) error {
	data := "{" +
		"\"username\":" + "\"" + username + "\"," +
		"\"password\":" + "\"" + password + "\"" +
		"}"

	req, err := http.NewRequest("POST", "https://firerain.me/api/login", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := result{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return err
	}

	if r.Code == 0 {
		a.cookies = resp.Cookies()
		return nil
	} else {
		return errors.New(r.Message)
	}
}

func (a *Account) GetItem() ([]Item, error) {
	req, err := http.NewRequest("GET", "https://firerain.me/api/item", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for _, cookie := range a.cookies {
		req.AddCookie(cookie)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := result{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Code == 0 {
		a.cookies = resp.Cookies()
		return r.Items, nil
	} else {
		return nil, errors.New(r.Message)
	}
}
