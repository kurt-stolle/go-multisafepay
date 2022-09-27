<p align="center">
    <img src="https://camo.githubusercontent.com/517483ae0eaba9884f397e9af1c4adc7bbc231575ac66cc54292e00400edcd10/68747470733a2f2f7777772e6d756c7469736166657061792e636f6d2f66696c6561646d696e2f74656d706c6174652f696d672f6d756c7469736166657061792d6c6f676f2d69636f6e2e737667" width="400px" position="center">
</p>

# Go wrapper for MultiSafepay
[MultiSafepay](https://www.multisafepay.com/about-us/) is a Dutch payment service provider that supports a range of local and international payment methods. This package is a wrapper for the [MultiSafepay API](https://docs.multisafepay.com/api/), created to simplify and standardize the integration into Go projects.

This project is maintained by the MultiSafepay community. There are several wrappers in other programming languages, some maintained by the development team and some by the community. See the page [Plugin integration]

Don't hesitate to contact me for questions about this wrapper. I'm a professional developer and partner of MultiSafepay. I've integrated several webshops with this payment service provider.

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
