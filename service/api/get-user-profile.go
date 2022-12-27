package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"
	"wasaphoto-1849661/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that retrives all the necessary infos of a profile
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	requestedUser := ps.ByName("id")

	var followers []database.User
	var following []database.User
	var photos []database.Photo

	// Check if the requesting user is banned by the requested profile owner
	userBanned, err := rt.db.BannedUserCheck(User{IdUser: requestingUserId}.ToDatabase(),
		User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/userBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if the requested profile was banned by the requesting user. If it's true respond with partial content
	requestedProfileBanned, err := rt.db.BannedUserCheck(User{IdUser: requestedUser}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/requestedProfileBanned: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestedProfileBanned {
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	userExists, err := rt.db.CheckUser(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.CheckUser: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !userExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	followers, err = rt.db.GetFollowers(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowers: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	following, err = rt.db.GetFollowing(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowing: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photos, err = rt.db.GetPhotosList(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetPhotosList: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nickname, err := rt.db.GetNickname(User{IdUser: requestedUser}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetNickname: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(CompleteProfile{
		Name:      requestedUser,
		Nickname:  nickname,
		Followers: followers,
		Following: following,
		Posts:     photos,
	})

}
