package apiRoutesV1

import (
	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/microservices/authenticationAPI/apiHelpers"
)

func Register(ctx *gin.Context) {
	apiHelpers.ErrorJSON(ctx, 200, "OK")
}
