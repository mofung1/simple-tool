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
    success: async (loginRes) => {
      try {
        // 调用登录接口
        const { data } = await wxLogin(loginRes.code)
        if (!data?.token || !data?.userInfo?.avatar) {
          throw new Error('登录失败，请重试')
        }

        // 保存用户信息
        userStore.setToken(data.token)
        userStore.setUserInfo(data.userInfo)
        
        options.success?.()
      } catch (error: any) {
        const errorMsg = error?.data?.message || '登录失败，请重试'
        options.fail?.(errorMsg)
        uni.showToast({
          title: errorMsg,
          icon: 'none'
        })
      }
    },
    fail: () => {
      const errorMsg = '微信登录失败，请重试'
      options.fail?.(errorMsg)
      uni.showToast({
        title: errorMsg,
        icon: 'none'
      })
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
