import { wxLogin, updateUserInfo } from '@/service/user'
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
          userStore.setUserInfo(res.data.user)
          // 登录成功提示
          uni.showToast({
            title: '登录成功',
            icon: 'success'
          })
          options.success?.()
        } else {
          throw new Error(res.message || '登录失败')
        }
      } catch (error) {
        const errorMsg = error.message || '登录失败，请重试'
        uni.showToast({
          title: errorMsg,
          icon: 'none'
        })
        options.fail?.(errorMsg)
      }
    },
    fail: (error) => {
      const errorMsg = error.errMsg || '微信登录失败，请重试'
      uni.showToast({
        title: errorMsg,
        icon: 'none'
      })
      options.fail?.(errorMsg)
    },
    complete: () => {
      options.complete?.()
    }
  })
}

// 检查登录状态
export const checkLogin = () => {
  const userStore = useUserStore()
  return new Promise<void>((resolve, reject) => {
    if (userStore.isLoggedIn()) {
      resolve()
    } else {
      login({
        success: () => resolve(),
        fail: (error) => reject(error)
      })
    }
  })
}

// 获取用户头像
export const getUserProfile = () => {
  return new Promise<IUserInfo>((resolve, reject) => {
    uni.getUserProfile({
      desc: '用于完善用户资料',
      success: async (res) => {
        try {
          const { nickName: nickname, avatarUrl: avatar } = res.userInfo
          const userStore = useUserStore()
          // 更新到后端
          const updateRes = await updateUserInfo({ nickname, avatar })
          if (updateRes.code === 200 && updateRes.data) {
            // 更新本地存储
            userStore.setUserInfo(updateRes.data.user)
            resolve(updateRes.data.user)
          } else {
            reject(new Error(updateRes.message || '更新用户信息失败'))
          }
        } catch (error) {
          reject(error)
        }
      },
      fail: (error) => {
        reject(new Error(error.errMsg || '获取用户信息失败'))
      }
    })
  })
}

// 退出登录
export const logout = () => {
  const userStore = useUserStore()
  userStore.clearUserInfo()
}
