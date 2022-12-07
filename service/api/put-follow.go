package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to the followers list of another
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the new follower id from the request body
	var new_follower User
	err := json.NewDecoder(r.Body).Decode(&new_follower)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the id of the follower in the request is the same of bearer
	if new_follower.IdUser != extractBearer(r.Header.Get("Authorization")) {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(errors.New("id in request and authentication not consistent")).Error("put-follow: users trying to identify as someone else")
		return
	}

	// Add the new follower in the db via db function
	err = rt.db.FollowUser(
		User{IdUser: ps.ByName("follower_id")}.ToDatabase(),
		User{IdUser: ps.ByName("id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
