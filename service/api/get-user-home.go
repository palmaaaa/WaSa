package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"
	"wasaphoto-1849661/service/database"

	"github.com/julienschmidt/httprouter"
)

// This function retrieves all the photos of the people that the user is following
func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	identifier := extractBearer(r.Header.Get("Authorization"))

	// A user can only see his/her home
	valid := validateRequestingUser(ps.ByName("id"), identifier)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	followers, err := rt.db.GetFollowing(User{IdUser: identifier}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var photos []database.Photo
	for _, follower := range followers {

		followerPhoto, err := rt.db.GetPhotosList(
			User{IdUser: identifier}.ToDatabase(),
			User{IdUser: follower.IdUser}.ToDatabase())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for i, photo := range followerPhoto {
			if i >= database.PhotosPerUserHome {
				break
			}
			photos = append(photos, photo)
		}

	}

	w.WriteHeader(http.StatusOK)

	// Send the output to the user. Instead of giving null for no matches return and empty slice of photos. ( ontrollaerrore)
	_ = json.NewEncoder(w).Encode(photos)
}
