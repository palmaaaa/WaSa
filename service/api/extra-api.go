package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

type Banned struct {
	IsBanned bool
}

type UserExistence struct {
	UserExists bool
}

// Checks if a user has been banned by another one (the one in the query)
func (rt *_router) checkBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requesting user wasn't banned by the photo owner
	banned, err := rt.db.BannedUserCheck(
		User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: ps.ByName("id")}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("extra-api/checkBan/db.BannedUserCheck: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Banned{IsBanned: banned})

}

// Checks if a user exists
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("id")

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	exists, err := rt.db.CheckUser(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("extra-api/getUser/db.CheckUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(UserExistence{UserExists: exists})

}

// Gets a single photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, "./media/bru/photos/brand.png")

	/*
		requestingUserId := extractBearer(r.Header.Get("Authorization"))
		requestedUser := ps.ByName("id")
		requestedPhoto := ps.ByName("photo_id")

		// Check if the user is logged
		if isNotLogged(requestingUserId) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	*/
	w.WriteHeader(http.StatusOK)

}
