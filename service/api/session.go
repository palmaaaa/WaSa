package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		// ctx.Logger.WithError(err).Error("session: error parsing json (invalid format)")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INVALID_JSON_ERROR_MSG})
		return
	} else if !validIdentifier(user.IdUser) {
		// Here we validated the user identifier and we discovered that it's not valid.
		w.WriteHeader(http.StatusBadRequest)
		// ctx.Logger.WithError(err).Error("session: invalid identificator lenght")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: INVALID_IDENTIFIER_ERROR_MSG})
		return
	}

	// Create the user in the database.
	err = rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		// In this case, there's a sql error since the resource already exists and can't be inserted again.
		// The identifier is returned as expected.
		// remove: ctx.Logger.WithError(err).Error("User already exists")
		_ = json.NewEncoder(w).Encode(user)
		return
	}

	// Create user's directories locally
	createUserFolder(user.IdUser, ctx)
	// Send the output to the user.
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// Function that creates a new subdir for the specified user
func createUserFolder(identifier string, ctx reqcontext.RequestContext) error {
	dir, err := os.Getwd()
	if err != nil {
		ctx.Logger.WithError(err).Error("session/createUserFolder: error getting current wd for user dir creating")
		return err
	}
	// Create the path media/useridentifier/ inside the project dir
	path := filepath.Join(filepath.Join(dir, "media"), identifier)

	// To the previously created path add the "photos" subdir
	err = os.MkdirAll(filepath.Join(path, "photos"), os.ModePerm)
	if err != nil {
		ctx.Logger.WithError(err).Error("session/createUserFolder:: error creating directories for user")
		return err
	}
	return nil
}
