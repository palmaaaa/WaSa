package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasaphoto-1849661/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	identifier := ps.ByName("id")

	// Copy body content (comment sent by user) into comment (struct)
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed to decode request body json")
		return
	}

	// Check if the comment has a valid lenght (<=30)
	if len(comment.Comment) > 30 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: comment longer than 30 characters")
		return
	}

	// Convert the photo identifier from string to int64
	photo_id_64, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
		return
	}

	// Function call to db for comment creation
	commentId, err := rt.db.CommentPhoto(
		PhotoId{IdPhoto: photo_id_64}.ToDatabase(),
		User{IdUser: identifier}.ToDatabase(),
		comment.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment: failed to execute query for insertion")
		return
	}

	// The response body will contain the unique id of the comment
	err = json.NewEncoder(w).Encode(CommentId{IdComment: commentId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
		return
	}
	w.WriteHeader(http.StatusCreated)
}
