package multisafepay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TestAPIURL is used to access the testing API
const TestAPIURL = "https://testapi.multisafepay.com/v1/json"

// LiveAPIURL is used to access the 'normal' API
const LiveAPIURL = "https://api.multisafepay.com/v1/json"

// Client handles all interaction with the MultiSafePay API
type Client struct {
	url    string
	apiKey string
}

// NewClient creates a new client with a provided URL and API key
func NewClient(url, apiKey string) *Client {
	return &Client{
		url:    url,
		apiKey: apiKey,
	}
}

// Get will perform a GET-request to the API
func (c *Client) Get(route string) (*http.Response, error) {
	// Prepare HTTP request
	req, err := http.NewRequest("GET", c.url+route, nil)
	if err != nil {
		return nil, err
	}

	// There is a bug in the MultiSafePay API where "api_key" must be lowercase.
	// Setting the Header map directly is a workaround for this issue.
	req.Header["api_key"] = []string{c.apiKey}

	// Do the request over the default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Post will perform a POST-request to the API with JSON-encoded data
func (c *Client) Post(route string, data interface{}) (*http.Response, error) {
	// Encode data
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(data); err != nil {
		return nil, err
	}

	// Prepare HTTP request
	req, err := http.NewRequest("POST", c.url+route, &buffer)
	if err != nil {
		return nil, err
	}

	// There is a bug in the MultiSafePay API where "api_key" must be lowercase.
	// Setting the Header map directly is a workaround for this issue.
	req.Header["api_key"] = []string{c.apiKey}

	// Content-Type header ensures the API reads our data as JSON
	req.Header.Set("Content-Type", "application/json")

	// Do the request over the default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateOrder creates a new order
func (c *Client) CreateOrder(o *Order) (*PostOrderResponse, error) {
	// Perform request to API
	res, err := c.Post("/orders", o)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	data := &PostOrderResponse{}
	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

// GetOrder fetches an order
func (c *Client) GetOrder(orderID string) (*GetOrderResponse, error) {
	// Perform request to API
	res, err := c.Get(fmt.Sprintf("/orders/%s", orderID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	data := &GetOrderResponse{}
	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
