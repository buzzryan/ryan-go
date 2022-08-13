package presenter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostContentRequest struct {
	ContentTitle string `json:"content_title" binding:"required"`
	ContentBody  string `json:"content_body" binding:"required"`
	ContentPath  string `json:"content_path" binding:"required"`
	ReadScope    string `json:"read_scope"`
	ModifyScope  string `json:"modify_scope"`
}

type PostContentResponse struct {
	PostContentRequest
	ContentID string `json:"content_id"`
}

// modify 는 edit 와 delete 가 포함된 개념입니다.

func PostContent(c *gin.Context) {
	req := PostContentRequest{}
	err := c.Bind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
func EditContent(ctx *gin.Context) {

}
func DeleteContent(ctx *gin.Context) {

}
func GetContent(ctx *gin.Context) {

}
