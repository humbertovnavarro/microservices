package context

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB
var API *gin.Engine
