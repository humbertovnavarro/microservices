package apiRoutesV1

import "github.com/gin-gonic/gin"

func InjectSelf(engine *gin.Engine) (engineV1 *gin.Engine) {
	v1 := engine.Group("/v1")
	v1.GET("ping", Ping)
	return engine
}
