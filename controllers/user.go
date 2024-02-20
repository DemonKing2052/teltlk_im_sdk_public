package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/utils"
	"ImSdk/svc"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   // 用户ID，可以根据实际情况定义更多的用户信息
	User_card string // 用户名，可以根据实际情况定义更多的用户信息
}

func Version(c *gin.Context) {

	c.JSON(200, gin.H{"code": 200, "message": e.GetMsg(e.SUCCESS), "result": "ok"})
}

// 生成参数appid,nonce ,sign
func Authenticate(c *gin.Context) {

	var res struct {
		Md5str string `json:"md_5_str"`
		Nonce  string `json:"nonce"`
	}

	res.Nonce = utils.GenerateRandomStr()
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(svc.Ctx.Config.Token.Secret + res.Nonce))
	md5Hash := hex.EncodeToString(hash[:])
	res.Md5str = md5Hash

	c.JSON(200, gin.H{"code": 200, "message": e.GetMsg(e.SUCCESS), "result": res})
}
