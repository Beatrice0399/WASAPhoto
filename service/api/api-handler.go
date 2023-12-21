package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	//rt.router.GET("/me", rt.wrap(rt.getMyProfile))
	rt.router.POST("/session", rt.doLogin)
	rt.router.GET("/users/:myid", rt.getMyProfile)
	rt.router.POST("/users/:myid", rt.setMyUsername)
	rt.router.GET("/users/:myid/profile/:pid", rt.getUserProfile)
	rt.router.POST("/users/:myid/photos", rt.uploadPhoto)
	rt.router.POST("/users/:myid/followers/:pid", rt.followUser)
	rt.router.DELETE("/users/:myid/followers/:pid", rt.unfollowUser)

	rt.router.POST("/users/:myid/banned/:pid", rt.banUser)
	rt.router.DELETE("/users/:myid/banned/:pid", rt.unbanUser)
	rt.router.GET("/users/:myid/stream", rt.getMyStream)
	rt.router.POST("/users/:myid/photos/:phid/like/", rt.likePhoto)
	rt.router.DELETE("/users/:myid/photos/:phid/like/", rt.unlikePhoto)

	rt.router.GET("/users/:myid/photos/:phid/", rt.getPhotoComments)
	rt.router.POST("/users/:myid/photos/:phid/comment/", rt.commentPhoto)
	rt.router.DELETE("/users/:myid/photos/:phid/comment/:cid", rt.uncommentPhoto)
	rt.router.DELETE("/users/:myid/photo/:phid", rt.deletePhoto)

	//utilities
	/*
		rt.router.GET("/photo/me", rt.getPhotoUser)
		rt.router.GET("/allProfiles", rt.getProfiles)
		rt.router.GET("/allUsers", rt.getUsers)
		rt.router.GET("/Follow", rt.getFollows)
		rt.router.GET("/me/banned", rt.getBanned)
		rt.router.GET("/ban", rt.getTableBan)
		rt.router.GET("/comments", rt.getTableComment)
		rt.router.GET("/likes", rt.getTableLikes)
	*/
	return rt.router
}
