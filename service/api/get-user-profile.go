package api

import (
	"wasaphoto-1849661/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)
// This function makes a call to the database to retrive all the users matching the query.
func (rt *_router) getUsersQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	identificator := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	// Search the user in the database (with the query parameter as a filter).
	res,err := rt.db.SearchUser(identificator)
	if err != nil {
		// In this case, there's an error coming from the database. Return an empty json.
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}

	// Send the output to the user.
	_ = json.NewEncoder(w).Encode(res)
}


