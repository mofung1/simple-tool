import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { wxLogin as wxLoginApi } from '@/service/user'
import type { IUserInfo } from '@/types/user'

const initState: IUserInfo = {
  id: 0,
  nickname: '',
  avatar: '',
  token: '',
  openid: '',
  mobile: '',
  gender: 0,
  country: '',
  province: '',
  city: ''
}

export const useUserStore = defineStore(
  'user',
  () => {
    const userInfo = ref<IUserInfo>({ ...initState })

    // 设置用户信息
    const setUserInfo = (val: Partial<IUserInfo>) => {
      userInfo.value = { ...userInfo.value, ...val }
    }

    // 清除用户信息
    const clearUserInfo = () => {
      userInfo.value = { ...initState }
      uni.removeStorageSync('token')
      uni.removeStorageSync('user-store')
    }

    // 检查是否登录
    const isLogined = computed(() => !!userInfo.value.token)

    // 微信登录
    const wxLogin = () => {
      return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        uni.login({
          provider: 'weixin',
          success: async (loginRes) => {
            try {
              const res = await wxLoginApi(loginRes.code)
              if (res.code === 200) {
                const { token, userInfo: user } = res.data
                setUserInfo({ ...user, token })
                uni.setStorageSync('token', token)
                resolve(user)
              } else {
                reject(new Error(res.message))
              }
            } catch (error) {
              reject(error)
            }
          },
          fail: (err) => {
            reject(new Error('微信登录失败：' + (err.errMsg || '未知错误')))
          }
        })
        // #endif

        // #ifdef H5
        reject(new Error('请在微信小程序中使用'))
        // #endif
      })
    }

    return {
      userInfo,
      isLogined,
      setUserInfo,
      clearUserInfo,
      wxLogin
    }
  },
  {
    persist: {
      key: 'user-store'
    }
  }
)
