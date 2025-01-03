import { defineStore } from 'pinia'
import { ref } from 'vue'
import { wxLogin, updateUserInfo } from '@/service/user'

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

    // 检查是否登录
    const isLoggedIn = (): boolean => {
      return !!token.value && !!userInfo.value.id
    }

    // 清除用户信息
    const clearUserInfo = () => {
      token.value = ''
      userInfo.value = {}
      uni.removeStorageSync('token')
      uni.removeStorageSync('userInfo')
    }

    // 获取存储的信息
    const loadStorageInfo = () => {
      const storageToken = uni.getStorageSync('token')
      const storageUserInfo = uni.getStorageSync('userInfo')
      if (storageToken) token.value = storageToken
      if (storageUserInfo) userInfo.value = storageUserInfo
    }

    // 微信登录
    const login = async () => {
      try {
        // 获取微信登录code
        const loginRes = await uni.login({ provider: 'weixin' })
        
        // 调用后端登录接口
        const res = await wxLogin(loginRes.code)
        if (res.code === 200 && res.data) {
          setToken(res.data.token)
          setUserInfo(res.data.user)
          return res.data
        } else {
          clearUserInfo()
          throw new Error(res.message || '登录失败')
        }
      } catch (error: any) {
        clearUserInfo()
        throw new Error(error.errMsg || error.message || '登录失败')
      }
    }

    // 获取用户信息并更新
    const getUserProfile = async () => {
      try {
        const profileRes = await uni.getUserProfile({ desc: '用于完善用户资料' })
        const { nickName: nickname, avatarUrl: avatar } = profileRes.userInfo
        
        // 更新到后端
        const res = await updateUserInfo({ nickname, avatar })
        if (res.code === 200 && res.data) {
          setUserInfo(res.data.user)
          return res.data.user
        } else {
          throw new Error(res.message || '更新用户信息失败')
        }
      } catch (error: any) {
        // 用户取消不清空状态
        if (!error.errMsg?.includes('cancel')) {
          clearUserInfo()
        }
        throw error
      }
    }

    // 退出登录
    const logout = () => {
      clearUserInfo()
    }

    return {
      token,
      userInfo,
      setToken,
      setUserInfo,
      loadStorageInfo,
      clearUserInfo,
      isLoggedIn,
      login,
      getUserProfile,
      logout
    }
  },
  {
    persist: true
  }
)
