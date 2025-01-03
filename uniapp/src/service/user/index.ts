import { http } from '@/utils/http'

// 微信登录
export const wxLogin = (code: string) => {
  return http.post<ILoginResult>('/api/v1/login/mnp', { code })
}

// 获取用户信息
export const getUserInfo = () => {
  return http.get<ILoginResult>('/api/user/info')
}

// 更新用户信息
export const updateUserInfo = (data: Partial<IUserInfo>) => {
  return http.post('/api/user/update', data)
}
