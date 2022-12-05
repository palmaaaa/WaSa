package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putNickname(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("id"), extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get the new nickname from the request body
	var nick Nickname
	err := json.NewDecoder(r.Body).Decode(&nick)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the username with the db function
	err = rt.db.ModifyNickname(ps.ByName("id"), nick.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: error executing update query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
