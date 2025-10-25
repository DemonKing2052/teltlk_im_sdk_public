package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/protos"
	"ImSdk/svc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ClientLogReporting(c *gin.Context) {
	req := protos.ClientLogReportingReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	//
	req.ClientLog = now.Format("2006-01-02 15:04:05") + " : " + req.ClientLog + "\n"

	svc.Ctx.ClientLogReq.FhMutex.Lock()
	defer svc.Ctx.ClientLogReq.FhMutex.Unlock()

	_, err := svc.Ctx.ClientLogReq.Fh.Write([]byte(req.ClientLog))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS)})
	return
}
