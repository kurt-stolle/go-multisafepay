package multisafepay

import (
	"encoding/json"
	"testing"
)

func TestID_MarshalJSON(t *testing.T) {
	// Define mock data
	data := make(map[string]interface{})
	data["number_based_id"] = ID("2015")
	data["string_based_id"] = ID("1958bravo")

	// Marshal mock data
	enc, err := json.Marshal(data)
	if err != nil {
		t.Error("could not encode:", err)
		t.FailNow()

		return
	}

	// Assert validity
	const expectedResponseData = `{"number_based_id":"2015","string_based_id":"1958bravo"}`

	if encString := string(enc); encString != expectedResponseData {
		t.Errorf("marshaled json (%s) does not match expected result (%s)", encString, expectedResponseData)
		t.FailNow()

		return
	}
}

func TestID_UnmarshalJSON(t *testing.T) {
	// Define mock JSON
	const data = `{"number_based_id":2015,"string_based_id":"1958bravo"}`

	// Decode data
	var dec struct {
		NumberBasedID ID `json:"number_based_id"`
		StringBasedID ID `json:"string_based_id"`
	}
	if err := json.Unmarshal([]byte(data), &dec); err != nil {
		t.Error("could not decode:", err)
		t.FailNow()

		return
	}

	// Assert validity
	if dec.NumberBasedID != "2015" || dec.StringBasedID != "1958bravo" {
		t.Errorf("json (%s) could not be decoded, values found: %+v", data, dec)
		t.FailNow()

		return
	}
}
