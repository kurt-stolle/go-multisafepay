package multisafepay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// performRequests executes a http.Request and returns the body or an error
func (c *Client) performRequest(req *http.Request) (io.ReadCloser, error) {
	// There is a bug in the MultiSafePay API where "api_key" must be lowercase.
	// Setting the Header map directly is a workaround for this issue.
	req.Header["api_key"] = []string{c.apiKey}

	// Do the request over the default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check status, anything but 200 (OK) is reason to raise an error
	if res.StatusCode != http.StatusOK {
		// Find error in the response body
		var errData ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errData); err != nil {
			return nil, APIError{
				Method:  req.Method,
				URL:     req.RequestURI,
				Status:  res.Status,
				Message: "error could not be decoded: " + err.Error(),
			}
		}

		return nil, APIError{
			Method:  req.Method,
			URL:     req.RequestURI,
			Status:  res.Status,
			Message: fmt.Sprintf("%s (error %d)", errData.ErrorInfo, errData.ErrorCode),
		}
	}

	// Return the response body
	return res.Body, nil
}

// Get will perform a GET-request to the API
func (c *Client) Get(route string) (io.ReadCloser, error) {
	// Prepare HTTP request
	req, err := http.NewRequest("GET", c.url+route, nil)
	if err != nil {
		return nil, err
	}

	// Do the request over the default client
	return c.performRequest(req)
}

// Post will perform a POST-request to the API with JSON-encoded data
func (c *Client) Post(route string, data interface{}) (io.ReadCloser, error) {
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

	// Content-Type header ensures the API reads our data as JSON
	req.Header.Set("Content-Type", "application/json")

	return c.performRequest(req)
}

// CreateOrder creates a new order
func (c *Client) CreateOrder(o Order) (*PostOrderResponse, error) {
	// Perform request to API
	res, err := c.Post("/orders", o)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Decode JSON
	data := &PostOrderResponse{}
	if err := json.NewDecoder(res).Decode(data); err != nil {
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
	defer res.Close()

	// Decode JSON
	data := &GetOrderResponse{}
	if err := json.NewDecoder(res).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}
