package main

import "github.com/gin-gonic/gin"
import "github.com/api"

func main() {

	router := gin.Default()

	// 创建api路由组
	api_group := router.Group("/api_group")

	// 在api路由组下创建usr路由组
	usr := api_group.Group("/usr")

	// 在usr路由组下创建loginCheck路由
	usr.POST("/loginCheck", api.LoginCheckPost)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
