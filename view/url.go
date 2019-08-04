package view

import (
	"HFish/view/api"
	"HFish/view/dashboard"
	"HFish/view/fish"
	"HFish/view/mail"
	"HFish/view/setting"
	"github.com/gin-gonic/gin"
	"HFish/view/login"
	"HFish/utils/conf"
	"net/http"
)

// 解决跨域问题
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func LoadUrl(r *gin.Engine) {
	// 登录
	r.GET("/login", login.Html)
	r.POST("/login", login.Login)
	r.GET("/logout", login.Logout)

	// 仪表盘
	r.GET("/", login.Jump, dashboard.Html)
	r.GET("/dashboard", login.Jump, dashboard.Html)
	r.GET("/get/dashboard/data", login.Jump, dashboard.GetFishData)

	// 钓鱼列表
	r.GET("/fish", login.Jump, fish.Html)
	r.GET("/get/fish/list", login.Jump, fish.GetFishList)
	r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
	r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

	// 邮件群发
	r.GET("/mail", login.Jump, mail.Html)
	r.POST("/post/mail/sendEmail",mail.SendEmailToUsers)

	// 设置
	r.GET("/setting", login.Jump, setting.Html)
	r.GET("/get/setting/info", setting.GetSettingInfo)
	r.POST("/post/setting/update",setting.UpdateEmailInfo)
	r.POST("/post/setting/login", setting.UpdateLoginInfo)
	r.POST("/post/setting/alertOver", setting.UpdateAlertOverInfo)
	r.POST("/post/setting/pushBullet", setting.UpdatePushBulletInfo)
	r.POST("/post/setting/pushFangTang", setting.UpdatePushFangTangInfo)
	r.POST("/post/setting/checkSetting",setting.UpdateStatusSetting)
	// API 接口
	// WEB 上报钓鱼信息
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		r.Use(cors())

		apiUrl := conf.Get("api", "url")
		r.POST(apiUrl, api.ReportWeb)
	}
}
