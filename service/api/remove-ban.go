package api

import (
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the banned list of another
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerToken := extractBearer(r.Header.Get("Authorization"))
	pathId := ps.ByName("id")
	userToUnban := ps.ByName("banned_id")

	// Check the user's identity for the operation
	valid := validateRequestingUser(pathId, bearerToken)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users can't ban themselfes so this action shouldn't be possible. In order to avoid
	// making any useless operation terminate here the execution of the function
	if userToUnban == bearerToken {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Remove the follower in the db via db function
	err := rt.db.UnbanUser(
		User{IdUser: pathId}.ToDatabase(),
		User{IdUser: userToUnban}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-ban: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
