import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

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
  city: ''
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
      uni.removeStorageSync('token')
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
            setUserInfo({
              ...userInfo,
              token: token.value
            })
            resolve(userInfo)
          },
          fail: (err) => {
            reject(err)
          }
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
              // TODO: 调用后端登录接口
              // const res = await login(loginRes.code)
              // setUserInfo(res.data.userInfo)
              resolve(loginRes)
            } catch (err) {
              reject(err)
            }
          },
          fail: (err) => {
            reject(err)
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
      token,
      setUserInfo,
      clearUserInfo,
      isLogined,
      getUserProfile,
      wxLogin
    }
  },
  {
    persist: {
      key: 'user-store',
      paths: ['userInfo', 'token']
    }
  }
)
