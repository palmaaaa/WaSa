package api

import (
	"net/http"
	"strconv"
	"strings"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a photo (this includes comments and likes)
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth := extractBearer(r.Header.Get("Authorization"))
	photoToDelete := ps.ByName("photo_id")

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Parse the photo in the url by removing anything after the dot (including the dot itself)
	photoInt, err := strconv.ParseInt(strings.Split(photoToDelete, ".")[0], 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error converting photoId to int")
		return
	}

	// Initialize the photo
	photo := PhotoId{IdPhoto: photoInt}

	// Call to the db function to remove the photo
	err = rt.db.RemovePhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error coming from database")
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
