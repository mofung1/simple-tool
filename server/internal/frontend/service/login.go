package service

import (
	"gorm.io/gorm"
	"simple-tool/server/internal/frontend/response"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/jwt"
	"simple-tool/server/pkg/wechat"
	"time"
)

type LoginService struct {
}

// MnpLogin 小程序登录
func (s *LoginService) MnpLogin(code string, clientIp string) (*response.LoginResult, error) {
	// 调用微信接口获取openid
	mp := &wechat.MnpProgram{
		AppID:     global.Conf.Wechat.AppID,
		AppSecret: global.Conf.Wechat.AppSecret,
	}

	wxResp, err := mp.Code2Session(code)
	if err != nil {
		global.ZapLog.Error("微信登录接口调用失败")
		return nil, err
	}

	var user models.User
	var userAuth models.UserAuth

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 查找用户授权信息
		result := tx.Where("auth_type = ? AND auth_id = ?", "mnp", wxResp.OpenID).First(&userAuth)
		if result.Error == nil {
			// 用户已存在，更新登录信息
			if err := tx.First(&user, userAuth.UserId).Error; err != nil {
				global.ZapLog.Error("查询用户信息失败")
				return err
			}
			// 更新用户登录信息
			if err := tx.Model(&user).Updates(map[string]interface{}{
				"login_time": time.Now(),
				"login_ip":   clientIp,
			}).Error; err != nil {
				global.ZapLog.Error("更新用户登录信息失败")
				return err
			}
		} else {
			// 创建新用户
			user = models.User{
				Nickname:  wxResp.Nickname,
				Avatar:    wxResp.HeadImgUrl,
				LoginTime: time.Now(),
				LoginIp:   clientIp,
			}
			if err := tx.Create(&user).Error; err != nil {
				global.ZapLog.Error("创建用户失败")
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
				global.ZapLog.Error("创建用户授权信息失败")
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 生成token
	token, err := jwt.GenToken(user.ID, "mnp")
	if err != nil {
		global.ZapLog.Error("生成token失败")
		return nil, err
	}

	return &response.LoginResult{
		Token: token,
		User:  user,
	}, nil
}
