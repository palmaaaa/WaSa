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

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get the new nickname from the request body
	var photo_like User
	err := json.NewDecoder(r.Body).Decode(&photo_like)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error decoding json")
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
	err = rt.db.LikePhoto(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), photo_like.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
