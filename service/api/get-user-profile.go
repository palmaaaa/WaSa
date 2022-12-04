package api

import (
	"encoding/json"
	"net/http"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// This function makes a call to the database to retrive all the users matching the query.
func (rt *_router) getUsersQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	identificator := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	var u User
	u.IdUser = identificator

	// Search the user in the database (with the query parameter as a filter).
	res, err := rt.db.SearchUser(u.ToDatabase())
	if err != nil {
		// In this case, there's an error coming from the database. Return an empty json.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Database has encountered an error")
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}

	// Send the output to the user. Instead of giving null for no matches return and empty slice of Users.
	if len(res) == 0 {
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}
