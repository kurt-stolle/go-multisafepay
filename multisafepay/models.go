package multisafepay

import "time"

// Response from the API, contains a boolean to indicate success
type Response struct {
	Success bool `json:"success"`
}

// ErrorResponse contains information about errors as reported by the API
type ErrorResponse struct {
	Response

	Data      interface{} `json:"data,omitempty"`
	ErrorCode int         `json:"error_code"`
	ErrorInfo string      `json:"error_info"`
}

// PaymentOptions structure, see: https://docs.multisafepay.com/api/#payment-option-object
type PaymentOptions struct {
	NotificationURL string `json:"notification_url,omitempty"`
	RedirectURL     string `json:"redirect_url,omitempty"`
	CancelURL       string `json:"cancel_url,omitempty"`
	CloseWindow     string `json:"close_window,omitempty"`
}

// Customer structure, see: https://docs.multisafepay.com/api/#customer-object
type Customer struct {
	Locale      string `json:"locale,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Company     string `json:"company,omitempty"`
	Address1    string `json:"address1,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	ZIPCode     string `json:"zip_code,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Referrer    string `json:"referrer,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
}

// Order structure, see: https://docs.multisafepay.com/api/#orders
type Order struct {
	Type           string          `json:"type,omitempty"`
	OrderID        string          `json:"order_id,omitempty"`
	Gateway        string          `json:"gateway,omitempty"`
	Currency       string          `json:"currency,omitempty"`
	Amount         int             `json:"amount,omitempty"`
	Description    string          `json:"description,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	Customer       *Customer       `json:"customer,omitempty"`
	SecondChance   struct {
		SendEmail bool `json:"send_email"`
	} `json:"second_chance"`
}

// PostOrderResponse is a response to POST /orders: https://docs.multisafepay.com/api/#orders
type PostOrderResponse struct {
	Response

	Data struct {
		OrderID    int    `json:"order_id"`
		PaymentURL string `json:"payment_url,omitempty"`
	} `json:"data"`
}

// Cost is a cost as presented in: https://docs.multisafepay.com/api/#retrieve-an-order
type Cost struct {
	TransactionID int       `json:"transaction_id"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	Status        string    `json:"status"`
	Created       time.Time `json:"created"`
	Amount        float64   `json:"amount"`
}

// GetOrderResponse is a response to GET /orders/{order_id}: https://docs.multisafepay.com/api/#retrieve-an-order
type GetOrderResponse struct {
	Response

	Data struct {
		TransactionID       int                    `json:"transaction_id"`
		OrderID             string                 `json:"order_id"`
		Created             time.Time              `json:"created"`
		Currency            string                 `json:"currency"`
		Amount              int                    `json:"amount"`
		Description         string                 `json:"description"`
		AmountRefunded      int                    `json:"amount_refunded"`
		Status              string                 `json:"status,omitempty"`
		FinancialStatus     string                 `json:"financial_status"`
		Reason              string                 `json:"reason"`
		ReasonCode          string                 `json:"reason_code"`
		FastCheckout        string                 `json:"fastcheckout"`
		Modified            time.Time              `json:"modified"`
		Customer            *Customer              `json:"customer"`
		PaymentDetails      map[string]interface{} `json:"payment_details"`
		Costs               []Cost                 `json:"costs"`
		RelatedTransactions []struct {
			Amount        int       `json:"amount"`
			Costs         []Cost    `json:"costs"`
			Created       time.Time `json:"created"`
			Currency      string    `json:"currency"`
			Description   string    `json:"description"`
			Modified      time.Time `json:"modified"`
			Status        string    `json:"status"`
			TransactionID int       `json:"transaction_id"`
		} `json:"related_transactions"`
		PaymentMethods []map[string]interface{} `json:"payment_methods"`
	}
}
