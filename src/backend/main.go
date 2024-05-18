package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	// 创建api路由组
	api := router.Group("/api")

	// 在api路由组下创建usr路由组
	usr := api.Group("/usr")

	// 在usr路由组下创建loginCheck路由
	usr.POST("/loginCheck", LoginCheckPost)
}
