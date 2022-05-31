package middleware

import (
	"errors"
	"go-examples/course/gateway/go_gateway_demo/public"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if sessInfo, ok := session.Get(public.AdminSessionInfoKey).(string); !ok || sessInfo == "" {
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
