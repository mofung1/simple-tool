import { http } from '@/utils/http'
import type { ILoginResult } from '@/types/user'

interface UserInfo {
  nickname?: string
  avatar?: string
  gender?: number
  country?: string
  province?: string
  city?: string
}

// 微信登录
export const wxLogin = (data: {
  code: string
  nickname: string
  avatar: string
  gender: number
  country: string
  province: string
  city: string
}) => {
  return http.post<ILoginResult>('/api/v1/login/mnp', data)
}

// 更新用户信息
export const updateUserInfo = (data: UserInfo) => {
  return http.post<ILoginResult>('/api/v1/user/update', data)
}
