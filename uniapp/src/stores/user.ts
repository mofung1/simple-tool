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
  city: '',
}

export const useUserStore = defineStore(
  'user',
  () => {
    const userInfo = ref<IUserInfo>({ ...initState })

    // 设置用户信息
    const setUserInfo = (val: Partial<IUserInfo>) => {
      userInfo.value = { ...userInfo.value, ...val }
      if (val.token) {
        uni.setStorageSync('token', val.token)
      }
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
    const wxLogin = async (params: { code: string; userInfo: WechatMiniprogram.UserInfo }) => {
      const res = await wxLoginApi({
        code: params.code,
        nickname: params.userInfo.nickName,
        avatar: params.userInfo.avatarUrl,
        gender: params.userInfo.gender,
        country: params.userInfo.country,
        province: params.userInfo.province,
        city: params.userInfo.city,
      })

      if (res.code === 200) {
        const { token, userInfo: user } = res.data
        setUserInfo({ ...user, token })
        return user
      }
      throw new Error(res.message || '登录失败')
    }

    // 处理登录态失效
    const handleTokenExpired = () => {
      clearUserInfo()
      // 跳转到登录页或刷新当前页面
      const pages = getCurrentPages()
      const currentPage = pages[pages.length - 1]
      if (currentPage?.route !== 'pages/user/user') {
        uni.reLaunch({ url: '/pages/user/user' })
      }
    }

    return {
      userInfo,
      isLogined,
      setUserInfo,
      clearUserInfo,
      wxLogin,
      handleTokenExpired,
    }
  },
  {
    persist: {
      key: 'user-store',
    },
  },
)
