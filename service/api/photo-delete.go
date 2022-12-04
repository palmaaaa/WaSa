package api

import (
	"net/http"
	"strconv"
	"strings"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth := extractBearer(r.Header.Get("Authorization"))
	photoToDelete := ps.ByName("photo_id")

	// If the user requesting to upload the photo has an invalid token then respond with a fobidden status.
	if auth == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// If the user requesting to upload the photo is not the same of the indicated user in the path then respond with an unathorized status.
	if ps.ByName("id") != auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the photo in the url by removing anything after the dot (including the dot itself)
	photoInt, err := strconv.ParseInt(strings.Split(photoToDelete, ".")[0], 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error converting photoId to int")
		return
	}
	var photo PhotoId
	photo.IdPhoto = photoInt

	err = rt.db.RemovePhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error coming from database")
	}
	w.WriteHeader(http.StatusNoContent)
}
