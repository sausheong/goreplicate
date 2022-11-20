package goreplicate

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var url = "https://api.replicate.com/v1/predictions"

// Create a new client
// Client is the main struct used to interact with the Replicate HTTP APIs
func NewClient(auth string, model *Model) *Client {
	return &Client{
		Authorization: auth,
		Model:         model,
		Request:       model.CreateRequest(),
	}
}

// Create a prediction
func (c *Client) Create() (err error) {
	// unmarshal the request to send to the API
	body, err := json.Marshal(c.Request)
	if err != nil {
		return
	}
	// create a HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a HTTP client and use it to send the request
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// read body from HTTP response and store into the client
	body, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &c.Response)
	return
}

// Get a prediction
func (c *Client) Get(predictionId string) (err error) {
	// create a HTTP GET request
	req, err := http.NewRequest("GET", url+"/"+predictionId, nil)
	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Content-Type", "application/json")

	// create a HTTP client and use it to send the request
	httpClient := http.Client{}

	ok := make(chan bool)
	go func() {
		for {
			// send a HTTP GET request to the API
			resp, err := httpClient.Do(req)
			if err != nil {
				break
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				break
			}
			err = json.Unmarshal(body, &c.Response)
			if err != nil {
				break
			}
			// TODO: check other states
			if c.Response.Status == "succeeded" {
				ok <- true
				break
			}
			if c.Response.Status == "failed" {
				ok <- true
				break
			}
		}
	}()
	<-ok
	return
}

// TODO: Get list of predictions
func (c *Client) List() (err error) {
	return
}

// TODO: Cancel a prediction
func (c *Client) Cancel(predictionId string) (err error) {
	return
}

// TODO: Get a model
func (c *Client) GetModel() (model Model, err error) {
	return
}

// TODO: Get list of model versions
func (c *Client) ListModelVersions(owner, name string) (err error) {
	return
}

// TODO: Get a model version
func (c *Client) GetModelVersion(owner, name, version string) (err error) {
	return
}

// TODO: Get a collection of models
func (c *Client) ListModels() (models []Model, err error) {
	return
}
