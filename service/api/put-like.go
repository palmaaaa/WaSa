package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that add a like of a user to a photo
func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoAuthor := ps.ByName("id")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// User is trying to like his/her photo
	if photoAuthor == requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
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

	// Get the new nickname from the request body
	var photo_like User
	err = json.NewDecoder(r.Body).Decode(&photo_like)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Sent identifier is not consistent with requestin user bearer id
	if photo_like.IdUser != requestingUserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert the path parameters photo_id from string to int64
	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error converting path param photo_id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the like in the db via db function
	err = rt.db.LikePhoto(
		PhotoId{IdPhoto: photo_id_64}.ToDatabase(),
		photo_like.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
