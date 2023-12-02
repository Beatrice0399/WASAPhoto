package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	//rt.router.GET("/me", rt.wrap(rt.getMyProfile))
	rt.router.POST("/session", rt.doLogin)
	rt.router.POST("/profile/:pid/followUser", rt.followUser)
	rt.router.POST("/profile/:pid/unfollowUser", rt.unfollowUser)
	rt.router.POST("/me/setMyUsername", rt.setMyUsername)
	rt.router.POST("/me/banned/:pid", rt.banUser)
	rt.router.DELETE("/me/banned/:pid", rt.unbanUser)

	//utilities
	rt.router.GET("/allProfiles", rt.getProfiles)
	rt.router.GET("/allUsers", rt.getUsers)
	rt.router.GET("/Follow", rt.getFollows)
	rt.router.GET("/me/banned", rt.getBanned)
	rt.router.GET("/ban", rt.getTableBan)
	return rt.router
}
