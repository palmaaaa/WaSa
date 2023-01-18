package api

import (
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to banned list of another
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pathId := ps.ByName("id")
	pathBannedId := ps.ByName("banned_id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation (only owner of the account can add a banned user to that account list)
	valid := validateRequestingUser(pathId, requestingUserId)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Check if the user is trying to ban himself/herself
	if requestingUserId == pathBannedId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the new banned user in the db via db function
	err := rt.db.BanUser(
		User{IdUser: pathId}.ToDatabase(),
		User{IdUser: pathBannedId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.BanUser: error executing insert query")

		// Something  didn't work internally
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ban implies removing the follow (if exists)
	err = rt.db.UnfollowUser(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: pathBannedId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser1: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// The banned user will not follow the user anymore or else will have the banner in his home
	err = rt.db.UnfollowUser(
		User{IdUser: pathBannedId}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser2: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
