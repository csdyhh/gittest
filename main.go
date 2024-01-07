package main

import (
	"awesomeProject/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret1111"))
	//配置session的中间件，store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	setupRoutes(r)

	r.Run(":8080")
}

func setupRoutes(r *gin.Engine) {
	employeeRouter := r.Group("/admin")
	{
		routers.SetupEmployeeRoutes(employeeRouter)
	}
}
