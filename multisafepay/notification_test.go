package multisafepay

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"testing"
	"time"
)

func TestValidatePostNotification(t *testing.T) {
	apiKey := "apiKey"
	payload := "payload"
	timestamp := []byte(strconv.FormatInt(time.Now().UTC().UnixNano(), 10))

	t.Logf("API Key: %s, payload: %s, timestamp: %s", apiKey, payload, timestamp)

	auth := &bytes.Buffer{}
	auth.Write(timestamp)
	auth.Write(postNotificationAuthSep)
	auth.Write(makeHMAC(timestamp, payload, apiKey).Sum(nil))
	authHeader := base64.StdEncoding.EncodeToString(auth.Bytes())

	t.Logf("Auth header: %s", authHeader)

	result, err := ValidatePostNotification(payload, authHeader, apiKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Result timestamp: %v", result)
}
