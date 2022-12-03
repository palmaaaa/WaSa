package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Error parsing json (invalid format)")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Invalid json format"})
		return
	} else if !valid_identifier(user.IdUser) {
		// Here we validated the user identifier and we discovered that it's not valid.
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Invalid identificator lenght")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Identifier must be a string between 3 and 16 characters"})
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the user in the database.
	err = rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		// In this case, there's a sql error since the resource already exists and can't be inserted again.
		// The identifier is returned as expected.
		ctx.Logger.WithError(err).Error("User already exists")
		_ = json.NewEncoder(w).Encode(user)
		return
	}

	// Send the output to the user.
	_ = json.NewEncoder(w).Encode(user)
}
