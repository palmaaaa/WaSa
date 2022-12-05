package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
		ctx.Logger.WithError(err).Error("remove-like: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert the path parameters photo_id from string to int64
	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like: error converting path param photo_id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the like in the db via db function
	err = rt.db.UnlikePhoto(PhotoId{IdPhoto: photo_id_64}.ToDatabase(), photo_like.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
