package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), extractBearer(r.Header.Get("Authentication")))
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
	err = rt.db.FollowUser(
		User{IdUser: ps.ByName("id")}.ToDatabase(),
		banned.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
