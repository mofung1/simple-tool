package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/jwt"
	"simple-tool/server/pkg/wechat"
	"time"
)

type Login struct {
}

type MnpLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// MnpLogin 小程序登录
func (l *Login) MnpLogin(c *gin.Context) {
	var req MnpLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, "参数错误")
		return
	}

	// 调用微信接口获取openid
	mp := &wechat.MnpConfig{
		AppID:     global.Conf.Wechat.AppID,
		AppSecret: global.Conf.Wechat.AppSecret,
	}

	wxResp, err := mp.Code2Session(req.Code)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	var user models.User
	var userAuth models.UserAuth

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 查找用户授权信息
		result := tx.Where("auth_type = ? AND auth_id = ?", "mnp", wxResp.OpenID).First(&userAuth)
		if result.Error == nil {
			// 用户已存在，更新登录信息
			if err := tx.First(&user, userAuth.UserId).Error; err != nil {
				return err
			}
			// 更新用户登录信息
			if err := tx.Model(&user).Updates(map[string]interface{}{
				"is_login":        1,
				"login_time":      time.Now(),
				"heart_beat_time": time.Now(),
			}).Error; err != nil {
				return err
			}
		} else {
			// 创建新用户
			user = models.User{
				IsLogin:       1,
				LoginTime:     time.Now(),
				HeartBeatTime: time.Now(),
			}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}

			// 创建用户授权信息
			userAuth = models.UserAuth{
				UserId:   user.ID,
				AuthType: "mnp",
				AuthId:   wxResp.OpenID,
				UnionId:  wxResp.UnionID,
			}
			if err := tx.Create(&userAuth).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		response.FailWithMsg(c, "保存用户信息失败")
		return
	}

	// 生成token
	token, err := jwt.GenToken(user.ID, "mnp")
	if err != nil {
		response.FailWithMsg(c, "生成token失败")
		return
	}

	response.OkWithData(c, gin.H{
		"token": token,
		"user":  user,
	})
}
