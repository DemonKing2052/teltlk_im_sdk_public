package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/protos"
	"ImSdk/common/utils"
	"ImSdk/svc"
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

func GetClientLogReporting(c *gin.Context) {
	req := protos.GetClientLogReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)
	start := page.Offset
	end := start + page.PageSize

	stat, err := svc.Ctx.ClientLogReq.Fh.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error()})
		return
	}

	var (
		size            = stat.Size()
		step      int64 = 1
		buf             = make([]byte, 0, 1024)
		lineBufs  [][]byte
		lineCount int
	)

	// 从文件末尾向前读取
	for pos := size - 1; pos >= 0; pos -= step {
		_, err := svc.Ctx.ClientLogReq.Fh.Seek(pos, 0)
		if err != nil {
			break
		}

		b := make([]byte, 1)
		if _, err := svc.Ctx.ClientLogReq.Fh.Read(b); err != nil {
			break
		}

		if b[0] == '\n' {
			// 收到换行，保存当前行
			lineCount++
			lineCopy := make([]byte, len(buf))
			copy(lineCopy, buf)
			lineBufs = append(lineBufs, lineCopy)
			buf = buf[:0]

			// 达到读取上限直接停止
			if lineCount >= end {
				break
			}
		} else {
			buf = append(buf, b[0])
		}
	}

	// 最后一行（如果没有换行符结尾）
	if len(buf) > 0 {
		lineCount++
		lineCopy := make([]byte, len(buf))
		copy(lineCopy, buf)
		lineBufs = append(lineBufs, lineCopy)
	}

	// 倒序拼接
	lines := make([]string, 0, end-start)
	for i := start; i < end && i < len(lineBufs); i++ {
		// 因为是倒序读的，每一行反转一次字节顺序
		reverseBytes(lineBufs[i])
		lines = append(lines, string(lineBufs[i]))
	}

	resp := protos.CommonListResp{
		List:     lines,
		Page:     page.Page,
		PageSize: page.PageSize,
		Total:    0,
	}
	resp.Total, _ = countLines(svc.Ctx.ClientLogReq.Fh)
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": resp})
	return
}

func countLines(f *os.File) (int64, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(f)
	var count int64
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return count, err
	}
	return count, nil
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
