package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))
	rt.router.POST("/users/:id/photos", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/users", rt.wrap(rt.getUsersQuery))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
