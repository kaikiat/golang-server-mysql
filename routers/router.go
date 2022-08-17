package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/kaikiat/golang-server-mysql-template/docs"
	"github.com/kaikiat/golang-server-mysql-template/routers/api"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", api.Ping)
	r.GET("/tags", api.GetTags)
	r.POST("/tags", api.AddTag)
	return r
}
