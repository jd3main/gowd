package gowd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

func (client *Client) Login(username, password string) (err error) {
	fmt.Println("getting login token")
	loginToken, err := client.GetToken("login")
	if err != nil {
		return err
	}

	fmt.Println("login token = ", loginToken)

	fmt.Println("logging in")

	err = client.clientLogin(loginToken, username, password)
	if err != nil {
		fmt.Println("Login Failed!")
		fmt.Println(err)
		return err
	}

	fmt.Println("Successfully Logged In")

	return nil
}

func (client *Client) clientLogin(loginToken, username, password string) error {
	endpoint := client.ServerAddr + MwApiPath

	queryBody := url.Values{
		"format":         {"json"},
		"action":         {"clientlogin"},
		"loginreturnurl": {"http://127.0.0.1"},
		"logintoken":     {loginToken},
		"username":       {username},
		"password":       {password},
	}
	//fmt.Println("queryBody:", queryBody)

	resp, err := client.PostForm(endpoint, queryBody)
	if err != nil {
		return err
	}

	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Clientlogin struct {
			Status   string
			Username string
		}
		Err json.RawMessage `json:"error"`
	}

	_ = json.Unmarshal(jsonBody, &result)

	if result.Clientlogin.Status == "PASS" {
		return nil
	}

	return errors.New(string(result.Err))
}
