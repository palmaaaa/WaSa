package api

import (
	"net/http"
	"strconv"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a like from a photo
func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoAuthor := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// User trying to unlike his/her photo. Since it's not possibile to like it in the first
	// place it's useless. Return to avoid doing useless operations
	if photoAuthor == requestingUserId {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BannedUserCheck(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: photoAuthor}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		// User was banned by owner, can't post the comment
		w.WriteHeader(http.StatusForbidden)
		return
	}

	photoIdInt, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/ParseInt: error converting photo_id to int64")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the like in the db via db function
	err = rt.db.UnlikePhoto(
		PhotoId{IdPhoto: photoIdInt}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/db.UnlikePhoto: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
