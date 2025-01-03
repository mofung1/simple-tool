import { useUserStore } from '@/stores/user'

// 获取token
export function getToken(): string {
  return uni.getStorageSync('token') || ''
}

// 设置token
export function setToken(token: string) {
  return uni.setStorageSync('token', token)
}

// 移除token
export function removeToken() {
  return uni.removeStorageSync('token')
}

// 检查是否登录
export function checkLogin(): boolean {
  const userStore = useUserStore()
  if (!userStore.isLogined) {
    uni.showToast({
      title: '请先登录',
      icon: 'none'
    })
    return false
  }
  return true
}

// 退出登录
export function logout() {
  const userStore = useUserStore()
  userStore.clearUserInfo()
}
