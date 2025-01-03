import { wxLogin } from '@/service/user'
import { useUserStore } from '@/stores/user'

interface LoginOptions {
  success?: () => void
  fail?: (error: string) => void
  complete?: () => void
}

// 微信登录
export const login = (options: LoginOptions = {}) => {
  const userStore = useUserStore()
  
  // 如果已经登录，直接返回
  if (userStore.isLoggedIn()) {
    options.success?.()
    options.complete?.()
    return
  }

  // 获取登录code
  uni.login({
    provider: 'weixin',
    success: async ({ code }) => {
      try {
        // 调用后端登录接口
        const res = await wxLogin(code)
        if (res.code === 200 && res.data) {
          // 保存token和用户信息
          userStore.setToken(res.data.token)
          userStore.setUserInfo(res.data.userInfo)
          options.success?.()
        } else {
          options.fail?.('登录失败，请重试')
        }
      } catch (error) {
        options.fail?.(error.message || '登录失败，请重试')
      }
    },
    fail: (error) => {
      options.fail?.(error.errMsg || '微信登录失败，请重试')
    },
    complete: () => {
      options.complete?.()
    }
  })
}

// 检查登录状态
export const checkLogin = () => {
  const userStore = useUserStore()
  return userStore.isLoggedIn()
}

// 获取用户信息
export const getUserInfo = () => {
  const userStore = useUserStore()
  return userStore.userInfo
}

// 退出登录
export const logout = () => {
  const userStore = useUserStore()
  userStore.clearUserInfo()
}
