package service

import (
	"fmt"
	"gorm.io/gorm"
	"simple-tool/server/internal/frontend/request"
	"simple-tool/server/internal/frontend/response"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/jwt"
	"simple-tool/server/pkg/utils"
	"simple-tool/server/pkg/wechat"
	"time"
)

type LoginService struct {
}

// MnpLogin 小程序登录
func (s *LoginService) MnpLogin(req request.MnpLoginRequest, clientIp string) (*response.LoginResult, error) {
	// 调用微信接口获取openid
	mp := &wechat.MnpProgram{
		AppID:     global.Conf.Wechat.AppID,
		AppSecret: global.Conf.Wechat.AppSecret,
	}

	wxResp, err := mp.Code2Session(req.Code)
	if err != nil {
		global.ZapLog.Error("微信登录接口调用失败")
		return nil, err
	}

	var user models.User
	var userAuth models.UserAuth

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 查找用户授权信息
		result := tx.Where("auth_type = ? AND openid = ?", "mnp", wxResp.OpenID).First(&userAuth)
		if result.RowsAffected > 0 {
			// 用户已存在，更新登录信息
			if err := tx.First(&user, userAuth.UserId).Error; err != nil {
				global.ZapLog.Error("查询用户信息失败")
				return err
			}
			// 更新用户登录信息
			if err := tx.Model(&user).Updates(map[string]interface{}{
				"avatar":     req.Avatar,
				"gender":     req.Gender,
				"login_time": time.Now(),
				"login_ip":   clientIp,
			}).Error; err != nil {
				global.ZapLog.Error("更新用户登录信息失败")
				return err
			}
		} else {
			sn := getUniqueUserSn()
			nickname := req.Nickname

			if req.Nickname == "微信用户" {
				nickname = nickname + fmt.Sprintf("%d", sn)
			}

			// 创建新用户
			user = models.User{
				Nickname:  nickname,
				Sn:        sn,
				Avatar:    req.Avatar,
				Gender:    req.Gender,
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
				Openid:   wxResp.OpenID,
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
		User: response.UserInfo{
			Id:       user.ID,
			Sn:       user.Sn,
			Avatar:   user.Avatar,
			Nickname: user.Nickname,
			Gender:   user.Gender,
		},
	}, nil
}

func getUniqueUserSn() int64 {
	var user = models.User{}
	sn := utils.GenRandSn()
	for {
		// 查询数据库中是否已存在该用户编号
		if tx := global.DB.Where("sn = ?", sn).First(&user); tx.RowsAffected == 0 {
			return sn
		}
		sn = utils.GenRandSn()
	}
}
