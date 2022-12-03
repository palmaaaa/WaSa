package api

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that manages the upload of a photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	auth := r.Header.Get("Authorization")
	w.Header().Set("Content-Type", "application/json")

	_, errPng := png.Decode(r.Body)
	_, errJpg := jpeg.Decode(r.Body)
	if errPng != nil && errJpg != nil {
		//fmt.Println("L'immagine deve essere nel formato jpeg o png")
		w.WriteHeader(http.StatusBadRequest)
		//ctx.Logger.WithError(errPng).Error("Error copying body content into file photo")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Images must be jpeg or png"})
		return
	}

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("errore")
		return
	}

	fmt.Printf("%s/%s/photos", curDir, auth)
	create_photo_folder(auth, curDir)
	// Generate a unique id for the photo

	// Create an empty file for storing the body content (image)
	out, err := os.Create("./Prova.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error creating local photo file")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

	// Copy body content to the previously created file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error copying body content into file photo")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INTERNAL_ERROR_MSG})
		return
	}

}
