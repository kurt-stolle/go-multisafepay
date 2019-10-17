# Go wrapper for MultiSafePay

This project is a wrapper for the [MultiSafePay API](https://docs.multisafepay.com/api/) in Go. 

```go
import "github.com/kurt-stolle/go-multisafepay/multisafepay"
```

## Documentation
See [GoDoc](https://godoc.org/github.com/kurt-stolle/go-multisafepay/multisafepay)

## Contributing
Pull requests and issues are welcome.

## Basic example
The following code example shows the creation of an order, and outputs the response from the API.
```go
package main

import (
  "errors"
  "fmt"
  
  "github.com/kurt-stolle/go-multisafepay/multisafepay"
)

func main() {
  // Set-up the client with the required parameters (api url and key)
  var client = multisafepay.NewClient(multisafepay.TestAPIURL, "my_api_key")
  
  // Define the order
  var order = multisafepay.Order{
    // order parameters go here, see documentation
  }
  
  // Create the order using the client defined above
  var response, err = client.CreateOrder(order)
  if err != nil {
    panic("could not create order: " + err.Error())
  }
  
  // Print the response data
  fmt.Print(response.Data)
}

```
