package trading212

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Login as a user obtaining the session cookies required to make other
// requests.
func (c *Client) Login(username, password string) error {
	token, err := c.getAuthToken()
	if err != nil {
		return fmt.Errorf("get auth token: %w", err)
	}

	reqURL := "https://www.trading212.com/en/authenticate"
	vs := url.Values{
		"login[username]":   {username},
		"login[password]":   {password},
		"login[_token]":     {token},
		"login[rememberMe]": {"1"},
	}
	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(vs.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (c *Client) getAuthToken() (string, error) {
	loginURL := "https://www.trading212.com/en/login"
	req, err := http.NewRequest(http.MethodGet, loginURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // nolint: errcheck

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse doc: %w", err)
	}

	token, ok := doc.Find(`input[name="login[_token]"]`).Attr("value")
	if !ok {
		return "", errors.New("login _token not found")
	}
	return token, nil
}
