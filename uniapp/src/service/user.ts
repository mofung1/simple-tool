import { http } from '@/utils/http'
import type { IUserInfo } from '@/types/user'

interface LoginParams {
  code: string
  nickname: string
  avatar: string
  gender: number
}

interface LoginResult {
  code: number
  message: string
  data: {
    token: string
    userInfo: IUserInfo
  }
}

// 微信登录
export const wxLogin = (data: LoginParams) => {
  return http.post<LoginResult>('/api/v1/user/login', data)
}

// 获取用户信息
export const getUserInfo = () => {
  return http.get<{
    code: number
    message: string
    data: IUserInfo
  }>('/api/user/info')
}

// 更新用户信息
export const updateUserInfo = (data: Partial<IUserInfo>) => {
  return http.post<{
    code: number
    message: string
    data: IUserInfo
  }>('/api/user/info', data)
}
