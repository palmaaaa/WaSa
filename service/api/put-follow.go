package api

import (
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"
	"wasaphoto-1849661/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to the followers list of another
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userToFollowId := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// users can't follow themselves
	if requestingUserId == userToFollowId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the id of the follower in the request is the same of the bearer and the path parameter
	if ps.ByName("follower_id") != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		// ctx.Logger.WithError(errors.New("id in request and authentication not consistent")).Error("put-follow: users trying to identify as someone else")
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BannedUserCheck(
		database.User{IdUser: requestingUserId},
		database.User{IdUser: userToFollowId})
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/rt.db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		// User was banned, can't perform the follow action
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Add the new follower in the db via db function
	err = rt.db.FollowUser(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: userToFollowId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
