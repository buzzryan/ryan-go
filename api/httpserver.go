package api

import "github.com/gin-gonic/gin"

func InitializeServerApp() *gin.Engine {
	engine := gin.Default()
	setRouter(engine)
	return engine
}

func setRouter(e *gin.Engine) {
	e.GET("")
	e.POST("")
	e.PUT("")
	e.DELETE("")
}
