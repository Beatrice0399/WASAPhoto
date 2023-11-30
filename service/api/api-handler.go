package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	//rt.router.GET("/me", rt.wrap(rt.getMyProfile))
	rt.router.POST("/session", rt.doLogin)
	rt.router.POST("/profile/:pid/followUser", rt.followUser)

	//utilities
	rt.router.GET("/allProfiles", rt.getProfiles)
	rt.router.GET("/allUsers", rt.getUsers)
	rt.router.GET("/Follow", rt.getFollows)
	return rt.router
}
