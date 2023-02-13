package jwt

import (
	"github.com/gin-gonic/gin"
	"soft-pro/resp"
)

// User类型用户权限校验
func UserJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 token
		token := context.GetHeader("token")

		if token == "" {
			resp.FailWithMessage(resp.TokenWithoutErrorMsg, context)
			context.Abort()
			return
		} else {
			// 解析 token
			claim, err := ParseToken(token)
			if err != nil {
				resp.FailWithMessage(err.Error(), context)
				context.Abort()
				return
			}
			// 校验缓存 token
			u, err := CheckBufferToken(token, claim.UserID)
			if err != nil {
				resp.FailWithMessage(err.Error(), context)
				context.Abort()
				return
			}
			// 通过校验
			context.Set("user", u)
			context.Next()
		}
	}
}

// Admin类型用户权限校验
func AdminJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 token
		token := context.GetHeader("token")

		if token == "" {
			resp.FailWithMessage(resp.TokenWithoutErrorMsg, context)
			context.Abort()
			return
		} else {
			// 解析 token
			claim, err := ParseToken(token)
			if err != nil {
				resp.FailWithMessage(err.Error(), context)
				context.Abort()
				return
			}
			// 校验缓存 token
			u, err := CheckBufferToken(token, claim.UserID)
			if err != nil {
				resp.FailWithMessage(err.Error(), context)
				context.Abort()
				return
			}
			// 校验权限
			if u.Role != "admin" {
				resp.FailWithMessage(resp.ForbiddenMsg, context)
				context.Abort()
				return
			}
			// 通过校验
			context.Set("user", u)
			context.Next()
		}
	}
}
