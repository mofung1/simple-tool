import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore(
  'user',
  () => {
    const token = ref<string>('')
    const userInfo = ref<IUserInfo>({})

    // 设置token
    const setToken = (value: string) => {
      token.value = value
      uni.setStorageSync('token', value)
    }

    // 设置用户信息
    const setUserInfo = (info: IUserInfo) => {
      userInfo.value = info
      uni.setStorageSync('userInfo', info)
    }

    // 获取存储的信息
    const loadStorageInfo = () => {
      const storageToken = uni.getStorageSync('token')
      const storageUserInfo = uni.getStorageSync('userInfo')
      if (storageToken) token.value = storageToken
      if (storageUserInfo) userInfo.value = storageUserInfo
    }

    // 清除用户信息
    const clearUserInfo = () => {
      token.value = ''
      userInfo.value = {}
      uni.removeStorageSync('token')
      uni.removeStorageSync('userInfo')
    }

    // 检查是否登录
    const isLoggedIn = (): boolean => {
      return !!token.value && !!userInfo.value.id
    }

    return {
      token,
      userInfo,
      setToken,
      setUserInfo,
      loadStorageInfo,
      clearUserInfo,
      isLoggedIn
    }
  }
)
