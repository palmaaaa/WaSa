package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to banned list of another
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Check the user's identity for the operation (only owner of the account can add a banned user to that account list)
	valid := validateRequestingUser(ps.ByName("id"), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get the banned user id from the request body
	var banned User
	err := json.NewDecoder(r.Body).Decode(&banned)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the new banned user in the db via db function
	err = rt.db.BanUser(
		User{IdUser: ps.ByName("id")}.ToDatabase(),
		banned.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
