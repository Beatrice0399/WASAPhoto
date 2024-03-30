package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	rt.router.GET("/users", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:uid", rt.wrap(rt.getProfile))
	rt.router.PUT("/users/:uid", rt.wrap(rt.setMyUsername))

	rt.router.POST("/users/:uid/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:uid/photos/:phid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:uid/photos/:phid", rt.wrap(rt.getPhoto))

	rt.router.PUT("/users/:uid/followers/:fid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:uid/followers/:fid", rt.wrap(rt.unfollowUser))

	rt.router.PUT("/users/:uid/bannedUsers/:bid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:uid/bannedUsers/:bid", rt.wrap(rt.unbanUser))

	rt.router.GET("/users/:uid/home", rt.wrap(rt.getMyStream))

	rt.router.PUT("/users/:uid/photos/:phid/likes/:lid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:uid/photos/:phid/likes/:lid", rt.wrap(rt.unlikePhoto))

	rt.router.POST("/users/:uid/photos/:phid/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:uid/photos/:phid/comments/:cid", rt.wrap(rt.uncommentPhoto))

	//utilities

	rt.router.GET("/photo/me", rt.getPhotoUser)
	rt.router.GET("/allProfiles", rt.getProfiles)
	rt.router.GET("/allUsers", rt.getUsers)
	rt.router.GET("/Follow", rt.getFollows)
	rt.router.GET("/me/banned", rt.getBanned)
	rt.router.GET("/ban", rt.getTableBan)
	rt.router.GET("/comments", rt.getTableComment)
	rt.router.GET("/likes", rt.getTableLikes)

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
