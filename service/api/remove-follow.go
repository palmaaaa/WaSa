package api

import (
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"
	"wasaphoto-1849661/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the follower list of another
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	oldFollower := ps.ByName("follower_id")
	photoOwnerId := ps.ByName("id")

	// Check if the id of the follower in the path is the same of bearer (no impersonation)
	valid := validateRequestingUser(oldFollower, requestingUserId)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users can't follow themselfes so the unfollow won't do anything
	if photoOwnerId == requestingUserId {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BannedUserCheck(
		database.User{IdUser: requestingUserId},
		database.User{IdUser: photoOwnerId})
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

	// Remove the follower in the db via db function
	err = rt.db.UnfollowUser(
		User{IdUser: oldFollower}.ToDatabase(),
		User{IdUser: photoOwnerId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-follow: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
