package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPInterceptor : http请求拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		//验证登录token是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			// w.WriteHeader(http.StatusForbidden)
			// token校验失败则跳转到登录页面
			c.Abort()
			c.Redirect(http.StatusFound, "/static/view/signin.html")
			return
		}
		c.Next()
	}
}
