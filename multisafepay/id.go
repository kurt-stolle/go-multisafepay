package multisafepay

import (
	"errors"
	"fmt"
	"strconv"
)

// ID is a type for decoding IDs in JSON responses from MultiSafePay that do not obey a strict rule with regards to which type is returned.
// When the  API determines that the order ID string only contains numbers, it returns the value as a `number` type, instead of the `string` type.
// This type's value is guaranteed to be represented as a string in JSON.
type ID string

// MarshalJSON implements the json.Marshaler interface
func (id ID) MarshalJSON() ([]byte, error) {
	// Always return the value as a string, thus: `"` + value + `"`
	return append(append([]byte{'"'}, []byte(id)...),'"'), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (id *ID) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		// Data is a string, set the value of `id` to whatever is in-between the double quotes
		*id = ID(data[1:len(data)-1])
	} else if idInt, err := strconv.Atoi(string(data)); err == nil {
		// Data is a number, set the value to the stringified version
		*id = ID(strconv.Itoa(idInt))
	} else {
		// Data is of a type that cannot be interpreted as an ID, return an error
		return errors.New(fmt.Sprintf("id (value %s) could not be interpreted as a string or number", string(data)))
	}

	return nil
}
