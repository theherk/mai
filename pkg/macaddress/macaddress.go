// Package macaddress provides the api.
// This could easily provide the full type definitions for the more
// detailed responses. It could probably do better error handling too.
package macaddress

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// URL of the root of the API.
// This probably isn't necessary, but if the api were provided at any
// other root, this would make that much easier. e.g. testing an api.
var URL = "https://api.macaddress.io/v1"

// API provides methods for calling the API.
// This type of implementation containing the client makes test
// verification much cleaner.
type API struct {
	Client Client
	Key    string
}

// Client is a barebones interface which http.Client must implement.
// It only needs to be the subset of the concrete type that is used. Its
// purpose is to allow this client to be intercepted under test to avoid
// calling outside the application.
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// preflight checks verifies configuration.
func (api *API) preflight() error {
	if api.Key == "" {
		return errors.New("key empty; no call")
	}
	if api.Client == nil {
		api.Client = new(http.Client)
	}
	return nil
}

// Get is the only implemented verb.
// Currently, it takes no options beyond the search query. It gives the
// response without any handling beyond delineation between good
// responses and errors. Errors are otherwise communicated on stderr and
// left to the user to deal with... for now.
func (api API) Get(s string) (string, error) {
	if err := api.preflight(); err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Authentication-Token", api.Key)
	q := req.URL.Query()
	q.Add("search", s)
	req.URL.RawQuery = q.Encode()
	res, err := api.Client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
