package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// This function retrieves all the photos of the people that the user is following
func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))

	/*
		// If the user is not logged in then respond with a 403 http status
		if identifier == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	*/

	// A user can only see his/her home
	valid := validateRequestingUser(ps.ByName("id"), identifier)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get all the photos from from people followed by the requesting user
	photos, err := rt.db.GetStream(User{IdUser: identifier}.ToDatabase())
	if err != nil {
		// In this case, there's an error coming from the database. Return an empty json.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Database has encountered an error")
		// controllaerrore
		_ = json.NewEncoder(w).Encode([]Photo{})
		return
	}

	w.WriteHeader(http.StatusOK)

	// Send the output to the user. Instead of giving null for no matches return and empty slice of photos.
	if len(photos) == 0 {
		// controllaerrore
		_ = json.NewEncoder(w).Encode([]Photo{})
		return
	}
	// controllaerrore
	_ = json.NewEncoder(w).Encode(photos)
}
