package router

import (
	"github.com/gin-gonic/gin"
)

func setRoomRouter(r *gin.Engine, handler room.Handler, txMiddleware gin.HandlerFunc) {
	roomGroup := r.Group("/room")
	{
		roomGroup.POST("/create", txMiddleware, handler.CreateRoomHandler)
		roomGroup.GET("/findAddableUserList", handler.GetAddableUserListHandler)
		roomGroup.GET("/findRoomListOfUser", handler.GetRoomListHandler)
		roomGroup.GET("/findUserListOfRoom", handler.GetUserListOfRoomHandler)
		roomGroup.POST("/updateLastReadMsgIdx", handler.UpdateLastReadMsgIndexHandler)
	}
}
