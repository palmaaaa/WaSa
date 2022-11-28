package api

import (
	"wasaphoto-1849661/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !valid_identifier(user.IdUser) {
		// Here we validated the user identifier and we discovered that it's not valid.
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message:"Identifier must be a string between 3 and 16 characters"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the user in the database.
	err = rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		// In this case, there's a sql error since the resource already exists and can't be inserted again. 
		// The identifier is returned as expected.
		_ = json.NewEncoder(w).Encode(user)
		return
	}

	// Send the output to the user.
	_ = json.NewEncoder(w).Encode(user)
}
