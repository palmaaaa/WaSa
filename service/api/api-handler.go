package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login enpoint
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))

	// Search endpoint
	rt.router.GET("/users", rt.wrap(rt.getUsersQuery))

	// Update nickname Endpoint
	rt.router.PUT("/users/:id", rt.wrap(rt.putNickname))

	// Like endpoint
	rt.router.PUT("/users/:id/photos/:photo_id/:like_id", rt.wrap(rt.putLike))
	rt.router.DELETE("/users/:id/photos/:photo_id/:like_id", rt.wrap(rt.deleteLike))

	// Photo Endpoint
	rt.router.POST("/users/:id/photos", rt.wrap(rt.postPhoto))
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
