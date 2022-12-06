package api

import (
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerToken := extractBearer(r.Header.Get("Authentication"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), bearerToken)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Remove the follower in the db via db function
	err := rt.db.UnfollowUser(
		User{IdUser: ps.ByName("id")}.ToDatabase(),
		User{IdUser: ps.ByName("banned_id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-ban: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}