package api

import (
	"errors"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the follower list of another
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	old_follower := ps.ByName("follower_id")

	// Check if the id of the follower in the path is the same of bearer (no impersonation)
	if old_follower != extractBearer(r.Header.Get("Authorization")) {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(errors.New("follower id in path and authorization not consistent")).Error("remove-follow: users trying to identify as someone else")
		return
	}

	// Remove the follower in the db via db function
	err := rt.db.UnfollowUser(
		User{IdUser: old_follower}.ToDatabase(),
		User{IdUser: ps.ByName("id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-follow: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
