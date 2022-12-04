package api

import "strings"

// Function that verifies if the identifier of a user has the right lenght
func validIdentifier(identifier string) bool {
	return len(identifier) >= 3 && len(identifier) <= 16
}

// Function that extracts the bearer token from the Authorization header
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return tokens[1]
	}
	return ""
}
