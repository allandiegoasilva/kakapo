package router

import (
	"kakapo/internal/shared/reply"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(rGroup *gin.RouterGroup) {
	users := rGroup.Group("/users")
	{
		users.POST("", func(ctx *gin.Context) {
			data := reply.Reply[reply.Empty]{
				Code: reply.ValidationError,
				Data: &reply.Empty{},
			}

			ctx.JSON(200, data)
		})
	}
}
