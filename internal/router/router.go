package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(rGroup *gin.RouterGroup) {
	RegisterUserRouter(rGroup)
}
