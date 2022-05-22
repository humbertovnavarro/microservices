package apiEngine

import (
	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/microservices/authenticationAPI/apiRoutesV1"
)

func New() (engine *gin.Engine) {
	engine = gin.Default()
	apiRoutesV1.InjectSelf(engine)
	return engine
}
