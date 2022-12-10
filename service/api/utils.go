package api

import (
	"net/http"
	"strings"
)

// Function that verifies if the identifier of a user has the right lenght
func validIdentifier(identifier string) bool {
	return len(identifier) >= 3 && len(identifier) <= 16
}

// Function that extracts the bearer token from the Authorization header
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

// Function that checks if the requesting user has a valid token for the specified endpoint. Returns 0 if it's valid, the error (as a int, representing the http status) otherwise
func validateRequestingUser(identifier string, auth string) int {

	// If the requesting user has an invalid token then respond with a fobidden status
	if auth == "" {
		return http.StatusForbidden
	}

	//  If the requesting user's id is different than the one in the path then respond with a unathorized status.
	if identifier != auth {
		return http.StatusUnauthorized
	}
	return 0
}

// Function checks if two identficators are the same
func checkEquality(a string, b string) bool {
	return a == b
}
