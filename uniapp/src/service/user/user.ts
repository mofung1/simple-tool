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
  msg: string
  data: {
    token: string
    user: IUserInfo
  }
}

// 微信登录
export const wxLogin = (data: LoginParams) => {
  return http.post<LoginResult>('/api/v1/user/login', data)
}
