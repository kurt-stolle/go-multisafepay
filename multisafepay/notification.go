package multisafepay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

// Separator used in the decoded Auth header
var postNotificationAuthSep = ":"

// ValidatePostNotification validates POST-type notifications from MultiSafepay
// See: https://docs.multisafepay.com/developer/api/notification-url/#post-notification-example
func ValidatePostNotification(payload string, authHeader string, apiKey string) (string, error) {
	// Base64 decode the Auth header
	auth, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return "", fmt.Errorf("auth decode failed: %w", err)
	}

	// Split the decoded Auth header using the colon as separator
	splitAuth := strings.Split(string(auth), postNotificationAuthSep)
	if len(splitAuth) != 2 {
		return "", fmt.Errorf("auth contents invalid")
	}

	// First entry is the timestamp second is the hash
	timestamp := splitAuth[0]
	authHash := splitAuth[1]

	// Compute SHA512 hash using the website API key as HMAC
	check := makeHMAC([]byte(timestamp), payload, apiKey)

	if check != authHash {
		return "", fmt.Errorf("check rejected")
	}

	return timestamp, nil
}

func makeHMAC(timestamp []byte, payload string, apiKey string) string {
	check := hmac.New(sha512.New, []byte(apiKey))
	check.Write(timestamp)
	check.Write([]byte(postNotificationAuthSep))
	check.Write([]byte(payload))
	return hex.EncodeToString(check.Sum(nil))
}
