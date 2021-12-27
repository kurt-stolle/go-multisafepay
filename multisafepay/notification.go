package multisafepay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
)

// Separator used in the decoded Auth header
var postNotificationAuthSep = []byte(":")

// ValidatePostNotification validates POST-type notifications from MultiSafepay
// See: https://docs.multisafepay.com/developer/api/notification-url/#post-notification-example
func ValidatePostNotification(payload string, authHeader string, apiKey string) (string, error) {
	// Base64 decode the Auth header
	auth, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return "", fmt.Errorf("auth decode failed: %w", err)
	}

	// Split the decoded Auth header using the colon as separator
	splitAuth := bytes.Split(auth, postNotificationAuthSep)
	if len(splitAuth) != 2 {
		return "", fmt.Errorf("auth contents invalid")
	}

	// First entry is the timestamp second is the hash
	timestamp := splitAuth[0]
	authHash := splitAuth[1]

	// Compute SHA512 hash using the website API key as HMAC
	check := makeHMAC(timestamp, payload, apiKey)

	if !hmac.Equal(check.Sum(nil), authHash) {
		return "", fmt.Errorf("check rejected")
	}

	return string(timestamp), nil
}

func makeHMAC(timestamp []byte, payload string, apiKey string) hash.Hash {
	check := hmac.New(sha512.New, []byte(apiKey))
	check.Write(timestamp)
	check.Write(postNotificationAuthSep)
	check.Write([]byte(payload))
	return check
}
