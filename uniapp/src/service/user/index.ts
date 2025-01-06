import { http } from '@/utils/http'

interface UserInfo {
  nickName: string
  avatarUrl: string
  gender: number
  country: string
  province: string
  city: string
}

// 微信登录
export const wxLogin = (data: {
  code: string
  userInfo: UserInfo
}) => {
  return http.post<any>('/api/v1/login/mnp', {
    code: data.code,
    nickname: data.userInfo.nickName,
    avatar: data.userInfo.avatarUrl,
    gender: data.userInfo.gender,
    country: data.userInfo.country,
    province: data.userInfo.province,
    city: data.userInfo.city,
  })
}

// 更新用户信息
export const updateUserInfo = (data: Partial<UserInfo>) => {
  return http.post<any>('/api/v1/user/update', data)
}
