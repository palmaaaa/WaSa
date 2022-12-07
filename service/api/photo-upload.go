package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that manages the upload of a photo
func (rt *_router) postPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	auth := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Initialize photo struct
	photo := Photo{
		Comments: 0,
		Likes:    0,
		Owner:    auth,
		Date:     time.Now().UTC(),
	}

	// Create a copy of the body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// After reading the body we won't be able to read it again. We'll reassign a "fresh" io.ReadCloser to the body
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Check if the body content is either a png or a jpeg image
	extension, formatErr := checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if formatErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(formatErr).Error("photo-upload: body contains file that is neither jpg or png")
		// controllaerrore
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: IMG_FORMAT_ERROR_MSG})
		return
	}

	// Body has been read in the previous function so it's necessary to reassign a io.ReadCloser to it
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// Generate a unique id for the photo
	photoIdInt, err := rt.db.CreatePhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error executing db function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save the default path before switching to the photo path
	originalPath, err := os.Getwd()
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error getting current working directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoId := strconv.FormatInt(photoIdInt, 10)

	// Create the user's folder locally to save his/her images
	newPhotoPath, err := getUserPhotoFolder(auth, ctx)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error getting user's photo folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Change path to the photo folder
	err = os.Chdir(newPhotoPath)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error changing directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create an empty file for storing the body content (image)
	out, err := os.Create(photoId + extension)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error creating local photo file")
		//  = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Copy body content to the previously created file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error copying body content into file photo")
		// controllaerrore
		// _ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Close created file
	out.Close()

	// Switch back to the default path
	err = os.Chdir(originalPath)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error changing directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// controllaerrore
	_ = json.NewEncoder(w).Encode(PhotoId{IdPhoto: photoIdInt})

}

// Function checks if the format of the photo is png or jpeg. Returns the format extension and an error
func checkFormatPhoto(body io.ReadCloser, newReader io.ReadCloser, ctx reqcontext.RequestContext) (string, error) {

	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {

		body = newReader
		_, errPng := png.Decode(body)
		if errPng != nil {
			return "", errors.New(IMG_FORMAT_ERROR_MSG)
		}
		return ".png", nil
	}
	return ".jpg", nil
}

// Function that returns the path of the photo folder for a certain user
func getUserPhotoFolder(user_id string, ctx reqcontext.RequestContext) (string, error) {
	curPath, err := os.Getwd()
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload/getUserPhotoFolder: error getting current working directory")
		return "", err
	}

	// Path of the photo dir
	photoPath := filepath.Join(filepath.Join(curPath, "media"), filepath.Join(user_id, "photos"))
	// Change path to ./media/user_id/photos/
	err = os.Chdir(photoPath)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload/getUserPhotoFolder: error changing directory")
		return "", err
	}

	return photoPath, nil
}
