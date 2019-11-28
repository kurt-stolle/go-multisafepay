<p align="center">
  <img src="https://www.multisafepay.com/img/multisafepaylogo.svg" width="400px" position="center">
</p>

# Go wrapper for MultiSafePay
[MultiSafePay](https://www.multisafepay.com/about-us/) is a payment services provider that supports a supports a large number of international payment methods. This package is a wrapper for the [MultiSafePay API](https://docs.multisafepay.com/api/), created with the goal of simplifying and standardizing the integration into Go projects. 

This project is maintained by the MultiSafePay community. There exist a number of wrappers in other programming languages, some maintained by the development team and some by the community, see the page [Plugin integration](https://docs.multisafepay.com/integrations/).

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
