package api

import (
	"net/http"
	"path/filepath"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that serves the requested photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	/*
		cwd, err := os.Getwd()
		if err != nil {
			ctx.Logger.WithError(err).Error("getPhoto/os.Getwd: error getting the cwd")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	*/

	http.ServeFile(w, r,
		filepath.Join(photoFolder, ps.ByName("id"), "photos", ps.ByName("photo_id")))

}
