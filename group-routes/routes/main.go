package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	getRoutes()
	_ = router.Run(":8080")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addUserRouter(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
