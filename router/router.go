//创建Gin框架的路由系统

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	/*创建路由分组：把所有 /api/auth 开头的接口归为一组
	r.Group("/api/auth")：创建分组，前缀是 /api/auth
	auth：分组名，后续所有以 auth. 开头的接口，都会自动带上 /api/auth 前缀
	作用：简化路径写法，比如 auth.POST("/login") 等价于 r.POST("/api/auth/login")，代码更整洁*/
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})

		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Register Success",
			})
		})
	}
	return r
}
