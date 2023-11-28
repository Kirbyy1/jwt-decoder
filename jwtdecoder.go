package jwtdecoder

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// DecodeJWT decodes a JWT without verifying the signature and returns the claims as a JSON-encoded string.
func DecodeJWT(tokenString string) (string, error) {
	// Split the JWT into its three parts: header, payload, and signature
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("Invalid JWT format")
	}

	// Decode the payload (second part) of the JWT
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("Error decoding JWT payload: %v", err)
	}

	// Unmarshal the payload into a map
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", fmt.Errorf("Error unmarshalling JWT claims: %v", err)
	}

	// Marshal the claims into a JSON string
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("Error encoding JWT claims as JSON: %v", err)
	}

	return string(claimsJSON), nil
}
