package agent

import (
	"github.com/candbright/util/xgin"
	"github.com/gin-gonic/gin"
)

func restTest(context *gin.Context) {
	xgin.Ok(context, nil)
}
