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

// mock api (will be removed)

func AcceptFriendReq(context *gin.Context) {}
func RemoveFriend(context *gin.Context)    {}
func SendFriendReq(context *gin.Context)   {}
func GetUserInfo(context *gin.Context)     {}
func UnlikeContent(context *gin.Context)   {}
func LikeContent(context *gin.Context)     {}
func GetAllContent(context *gin.Context)   {}
func SignUp(ctx *gin.Context)              {}

func PostComment(ctx *gin.Context)   {}
func EditComment(ctx *gin.Context)   {}
func DeleteComment(ctx *gin.Context) {}
func GetComment(ctx *gin.Context)    {}
func UploadImage(ctx *gin.Context)   {}
