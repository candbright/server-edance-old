package main

import (
	"github.com/candbright/edance"
	"github.com/candbright/edance/agent"
	"github.com/candbright/util/xlog"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	edance.InitGlobal()
	xlog.InitLog(edance.LogFile)
	agent.InitManager()
	eng := gin.New()
	eng.Use(xlog.HandlerFunc("edance"))
	agent.RegisterMiddleware(eng)
	agent.RegisterHandlers(eng)
	xlog.Info("======================")
	xlog.Info("===     edance     ===")
	xlog.Info("======================")
	_ = eng.Run(":" + strconv.Itoa(edance.Port))
}
