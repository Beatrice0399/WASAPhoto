package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	rt.router.GET("/users/:myid", rt.wrap(rt.getMyProfile))
	rt.router.PUT("/users/:myid", rt.wrap(rt.setMyUsername))
	rt.router.GET("/users/:myid/profile/:pid", rt.wrap(rt.getUserProfile))

	rt.router.POST("/users/:myid/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:myid/photo/:phid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:myid/photos/:phid/", rt.wrap(rt.getPhotoComments))

	rt.router.PUT("/users/:myid/followers/:fid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:myid/followers/:fid", rt.wrap(rt.unfollowUser))

	rt.router.PUT("/users/:myid/banned/:bid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:myid/banned/:bid", rt.wrap(rt.unbanUser))

	rt.router.GET("/users/:myid/stream", rt.wrap(rt.getMyStream))

	rt.router.PUT("/users/:myid/photos/:phid/likes/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:myid/photos/:phid/likes/:lid", rt.wrap(rt.unlikePhoto))

	rt.router.POST("/users/:myid/photos/:phid/comment/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:myid/photos/:phid/comment/:cid", rt.wrap(rt.uncommentPhoto))

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
