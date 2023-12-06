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
	rt.router.GET("/profile/:pid", rt.getProfile)
	rt.router.POST("/me/photo", rt.uploadPhoto)
	rt.router.GET("/photo/me", rt.getPhotoUser)
	rt.router.GET("/photos/:phid/", rt.getPhotoComments)
	rt.router.POST("/photos/:phid/comment/", rt.commentPhoto)
	rt.router.DELETE("/photos/:phid/comment/:cid", rt.uncommentPhoto)
	rt.router.POST("/photos/:phid/like/", rt.likePhoto)
	rt.router.DELETE("/photos/:phid/like/", rt.unlikePhoto)
	rt.router.DELETE("/me/photo/:phid", rt.deletePhoto)
	rt.router.GET("/home", rt.getMyStream)

	//utilities
	rt.router.GET("/allProfiles", rt.getProfiles)
	rt.router.GET("/allUsers", rt.getUsers)
	rt.router.GET("/Follow", rt.getFollows)
	rt.router.GET("/me/banned", rt.getBanned)
	rt.router.GET("/ban", rt.getTableBan)
	rt.router.GET("/comments", rt.getTableComment)
	rt.router.GET("/likes", rt.getTableLikes)
	return rt.router
}
