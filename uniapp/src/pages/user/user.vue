<route lang="json5">
{
  style: {
    navigationBarTitleText: '',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-100">
    <!-- 用户信息卡片 -->
    <view class="bg-gradient-to-b from-blue-500 to-blue-600 p-6">
      <view v-if="userStore.isLogined" class="flex items-center space-x-4">
        <image
          :src="userStore.userInfo.avatar"
          class="w-16 h-16 bg-white rounded-full border-4 border-blue-400/30"
        />
        <view class="flex-1">
          <text class="text-white text-lg font-medium block mb-1">
            {{ userStore.userInfo.nickname }}
          </text>
        </view>
        <view class="i-carbon-chevron-right text-white"></view>
      </view>
      <button
        v-else
        open-type="getUserInfo"
        type="primary"
        class="flex items-center space-x-4 w-full bg-transparent border-none p-0"
        @getuserinfo="handleGetUserInfo"
      >
        <view class="w-16 h-16 bg-white rounded-full flex items-center justify-center shadow-lg">
          <view class="i-carbon-user-avatar text-blue-500 text-2xl"></view>
        </view>
        <view class="flex-1">
          <text class="text-white text-lg font-medium block mb-1">未登录</text>
          <text class="text-blue-100 text-sm">点击登录账号</text>
        </view>
        <view class="i-carbon-chevron-right text-white"></view>
      </button>
    </view>

    <!-- 功能区域 -->
    <view class="p-4 space-y-4">
      <!-- 常用功能 -->
      <view class="bg-white rounded-xl p-4 shadow-sm">
        <view class="grid grid-cols-4 gap-4">
          <view class="flex flex-col items-center">
            <view class="w-12 h-12 bg-blue-50 rounded-xl flex items-center justify-center mb-2">
              <view class="i-carbon-document text-blue-500 text-xl"></view>
            </view>
            <text class="text-sm text-gray-700">历史记录</text>
          </view>
          <view class="flex flex-col items-center">
            <view class="w-12 h-12 bg-purple-50 rounded-xl flex items-center justify-center mb-2">
              <view class="i-carbon-favorite text-purple-500 text-xl"></view>
            </view>
            <text class="text-sm text-gray-700">我的收藏</text>
          </view>
          <view class="flex flex-col items-center">
            <view class="w-12 h-12 bg-green-50 rounded-xl flex items-center justify-center mb-2">
              <view class="i-carbon-download text-green-500 text-xl"></view>
            </view>
            <text class="text-sm text-gray-700">下载记录</text>
          </view>
          <view class="flex flex-col items-center">
            <view class="w-12 h-12 bg-orange-50 rounded-xl flex items-center justify-center mb-2">
              <view class="i-carbon-share text-orange-500 text-xl"></view>
            </view>
            <text class="text-sm text-gray-700">分享应用</text>
          </view>
        </view>
      </view>

      <!-- 其他功能列表 -->
      <view class="bg-white rounded-xl shadow-sm overflow-hidden">
        <view class="divide-y divide-gray-100">
          <view class="flex items-center justify-between p-4" @tap="handleNavigation('help')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-help text-gray-400"></view>
              <text class="text-gray-700">常见问题</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view>
          <view class="flex items-center justify-between p-4" @tap="handleNavigation('about')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-information text-gray-400"></view>
              <text class="text-gray-700">关于我们</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

// 处理获取用户信息
const handleGetUserInfo = async (e: any) => {
  try {
    console.log('获取用户信息事件：', e)
    const userInfo = e.detail?.userInfo

    // 用户拒绝授权
    if (!userInfo) {
      uni.showToast({
        title: '需要您的授权才能继续',
        icon: 'none'
      })
      return
    }

    // 获取登录凭证
    const { code } = await uni.login({
      provider: 'weixin'
    })

    console.log('登录凭证：', code)
    console.log('用户信息：', userInfo)

    // 调用后端登录接口
    try {
      uni.showLoading({ title: '登录中...' })
      const user = await userStore.wxLogin({
        code,
        userInfo
      })
      uni.hideLoading()

      // 添加调试日志
      console.log('登录成功，用户信息：', user)
      console.log('store中的用户信息：', userStore.userInfo)
      console.log('是否登录：', userStore.isLogined)

      uni.showToast({
        title: '登录成功',
        icon: 'success',
      })
    } catch (error: any) {
      uni.hideLoading()
      console.error('登录失败：', error)
      console.error('错误堆栈：', error.stack)
      uni.showToast({
        title: error.message || '登录失败',
        icon: 'error',
      })
    }
  } catch (error: any) {
    console.error('获取用户信息失败：', error)
    console.error('错误堆栈：', error.stack)
    uni.showToast({
      title: '登录失败',
      icon: 'error',
    })
  }
}

// 处理页面导航
const handleNavigation = (type: string) => {
  switch (type) {
    case 'help':
      uni.navigateTo({
        url: '/pages/help/index',
      })
      break
    case 'about':
      uni.navigateTo({
        url: '/pages/about/index',
      })
      break
  }
}
</script>

<style lang="scss" scoped>
:deep(.wd-button) {
  border-radius: 0.5rem;
}
</style>
