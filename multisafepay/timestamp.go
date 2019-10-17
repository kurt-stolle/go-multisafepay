package multisafepay

import (
	"time"
)

// RFC3339NZ is time.RFC3339 without timezones
const RFC3339NZ = "2006-01-02T15:04:05"

// Time format used by MultiSafePay. RFC3339 without timezone
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(RFC3339NZ)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, RFC3339NZ)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	t.Time, err = time.Parse(`"`+RFC3339NZ+`"`, string(data))

	return err
}
