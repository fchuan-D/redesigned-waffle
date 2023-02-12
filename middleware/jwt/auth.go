package jwt

import (
	"github.com/gin-gonic/gin"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/resp"
)

func JWT() gin.HandlerFunc {
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
				resp.FailWithMessage(resp.TokenInValidErrorMsg, context)
				context.Abort()
				return
			} else {
				// 合法的 token,进一步校验是否过期
				// ToDo: 从 Redis读取是否有缓存的 User信息
				var u entity.User
				err := dao.Rd.Get(token).Scan(&u)
				// token未被缓存,已过期
				if err != nil {
					resp.FailWithMessage(resp.TokenExpiredErrorMsg, context)
					context.Abort()
					return
				} else if u.ID != claim.UserID {
					// 缓存数据出错
					resp.FailWithMessage(resp.InternalServerErrorMsg, context)
					context.Abort()
					return
				}
				// 通过校验
				context.Set("user", u)
				context.Next()
			}
		}
	}
}
