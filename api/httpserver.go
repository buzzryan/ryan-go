package api

import (
	. "github.com/buzzryan/ryan-go/api/presenter"
	"github.com/gin-gonic/gin"
)

func InitializeServerApp() *gin.Engine {
	engine := gin.Default()
	engine = setRouter(engine)
	return engine
}

func setRouter(e *gin.Engine) *gin.Engine {
	e.POST("/sign-in", SignIn)
	e.POST("/sign-out", SignOut)
	e.POST("/sign-up", SignUp)

	content := e.Group("/content")
	{
		content.GET("", GetContent)
		content.GET("", GetAllContent)
		content.POST("", PostContent)

		specificContent := content.Group("/:contentID")
		{
			specificContent.PUT("", EditContent)
			specificContent.DELETE("", DeleteContent)

			specificContent.PUT("/like", LikeContent)
			specificContent.PUT("/unlike", UnlikeContent)

			comment := content.Group("/comment")
			{
				comment.GET("", GetComment)
				comment.POST("", PostComment)
				comment.PUT("/:commentID", EditComment)
				comment.DELETE("/:commentID", DeleteComment)
			}
		}
	}

	e.POST("upload-image", UploadImage)

	user := content.Group("/user/:userID")
	{
		user.GET("", GetUserInfo)
		user.POST("/friend", SendFriendReq)
		user.POST("/friend", AcceptFriendReq)
		user.DELETE("/friend", RemoveFriend)
	}
	return e
}
