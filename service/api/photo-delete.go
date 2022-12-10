package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a photo (this includes comments and likes)
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth := extractBearer(r.Header.Get("Authorization"))
	photoIdStr := ps.ByName("photo_id")

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Parse the photo in the url by removing anything after the dot (including the dot itself)
	photoInt, err := strconv.ParseInt(photoIdStr, 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error converting photoId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Initialize the photo
	photo := PhotoId{IdPhoto: photoInt}

	// Call to the db function to remove the photo
	err = rt.db.RemovePhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete: error coming from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the folder of the file that has to be eliminated
	curPath, pathPhoto, err := getUserPhotoFolder(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/getUserPhotoFolder: error with directories")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove the file from the users' photos folder
	err = os.Remove(filepath.Join(pathPhoto, photoIdStr))
	if err != nil {
		// Error occurs if the file doesn't exist, but for idempotency an error won't be raised
		ctx.Logger.WithError(err).Error("photo-delete/os.Remove: photo to be removed is missing")
	}

	// Change the directory back to the previous path
	err = os.Chdir(curPath)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/os.Chdir: error changing directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
