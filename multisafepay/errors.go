package multisafepay

import "fmt"

// APIError reports error responses from API requests, i.e. when the status is anything other than 200
type APIError struct {
	Method  string
	URL     string
	Status  string
	Message string
}

// Error implements builtin error interface
func (e APIError) Error() string {
	return fmt.Sprintf("%s %s: %s: %s", e.Method, e.URL, e.Status, e.Message)
}
