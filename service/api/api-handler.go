package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	//rt.router.GET("/me", rt.wrap(rt.getMyProfile))
	rt.router.POST("/session", rt.doLogin)

	//utilities
	rt.router.GET("/allUsers", rt.getProfiles)
	return rt.router
}
