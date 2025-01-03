import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { wxLogin as wxLoginApi } from '@/service/user'

interface IUserInfo {
  nickname: string
  avatar: string
  token: string
  openid: string
  mobile: string
  gender: number
  country: string
  province: string
  city: string
}

const initState: IUserInfo = {
  nickname: '',
  avatar: '',
  token: '',
  openid: '',
  mobile: '',
  gender: 0,
  country: '',
  province: '',
  city: '',
}

export const useUserStore = defineStore(
  'user',
  () => {
    const userInfo = ref<IUserInfo>({ ...initState })
    const token = ref<string>('')

    // 设置用户信息
    const setUserInfo = (val: IUserInfo) => {
      userInfo.value = { ...userInfo.value, ...val }
      if (val.token) {
        token.value = val.token
      }
    }

    // 清除用户信息
    const clearUserInfo = () => {
      userInfo.value = { ...initState }
      token.value = ''
      // 清除本地存储
      uni.removeStorageSync('token')
      uni.removeStorageSync('user-store')
    }

    // 检查是否登录
    const isLogined = computed(() => !!token.value)

    // 获取用户信息
    const getUserProfile = () => {
      return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        uni.getUserProfile({
          desc: '用于完善用户资料',
          success: (res) => {
            const { userInfo } = res
            console.log('获取用户信息成功：', userInfo)
            // setUserInfo({
            //   ...userInfo,
            //   token: token.value,
            // })
            resolve(userInfo)
          },
          fail: (err) => {
            console.error('获取用户信息失败：', err)
            reject(new Error('获取用户信息失败：' + (err.errMsg || '未知错误')))
          },
        })
        // #endif

        // #ifdef H5
        // H5端可以实现其他登录方式
        reject(new Error('请在微信小程序中使用'))
        // #endif
      })
    }

    // 微信登录
    const wxLogin = () => {
      return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        uni.login({
          provider: 'weixin',
          success: async (loginRes) => {
            try {
              console.log('微信登录信息：', loginRes)
              // 调用后端登录接口
              const res = await wxLoginApi(loginRes.code)
              if (res.code === 200) {
                const { token, userInfo } = res.data
                setUserInfo({ ...userInfo, token })
                // 存储token
                uni.setStorageSync('token', token)
                resolve(userInfo)
              } else {
                throw new Error(res.message || '登录失败')
              }
            } catch (err) {
              console.error('登录失败：', err)
              reject(err)
            }
          },
          fail: (err) => {
            console.error('微信登录失败：', err)
            reject(new Error('微信登录失败：' + (err.errMsg || '未知错误')))
          },
        })
        // #endif

        // #ifdef H5
        reject(new Error('请在微信小程序中使用'))
        // #endif
      })
    }

    return {
      userInfo,
      token,
      setUserInfo,
      clearUserInfo,
      isLogined,
      getUserProfile,
      wxLogin,
    }
  },
  {
    persist: {
      key: 'user-store',
      paths: ['userInfo', 'token'],
    },
  },
)
