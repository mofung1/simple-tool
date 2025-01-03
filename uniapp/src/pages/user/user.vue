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
          <text class="text-blue-100 text-sm">{{ userStore.userInfo.mobile || '未绑定手机' }}</text>
        </view>
        <view class="i-carbon-chevron-right text-white"></view>
      </view>
      <view v-else class="flex items-center space-x-4" @click="handleLogin">
        <view class="w-16 h-16 bg-white rounded-full flex items-center justify-center shadow-lg">
          <view class="i-carbon-user-avatar text-blue-500 text-2xl"></view>
        </view>
        <view class="flex-1">
          <text class="text-white text-lg font-medium block mb-1">未登录</text>
          <text class="text-blue-100 text-sm">点击登录账号</text>
        </view>
        <view class="i-carbon-chevron-right text-white"></view>
      </view>
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
          <view class="flex items-center justify-between p-4" @click="handleNavigation('help')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-help text-gray-400"></view>
              <text class="text-gray-700">帮助与反馈</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view>
          <view class="flex items-center justify-between p-4" @click="handleNavigation('about')">
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

<script lang="ts" setup>
import { useUserStore } from '@/store/user'

const userStore = useUserStore()

// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

// 处理登录
const handleLogin = async () => {
  try {
    // 先获取用户信息（这个必须在用户点击时直接调用）
    const userInfo = await userStore.getUserProfile()

    uni.showLoading({
      title: '登录中...',
    })

    // 微信登录获取code
    const loginRes = await userStore.wxLogin()
    console.log('登录成功:', loginRes)
    // TODO: 调用后端登录接口
    const res = await login(loginRes.code)
    userStore.setUserInfo(res.data.userInfo)

    uni.hideLoading()
    uni.showToast({
      title: '登录成功',
      icon: 'success',
    })
  } catch (err: any) {
    uni.hideLoading()
    // 如果是用户取消，不显示错误提示
    if (err.errMsg && err.errMsg.includes('cancel')) {
      return
    }
    uni.showToast({
      title: '登录失败',
      icon: 'error',
    })
    console.error('登录失败:', err)
  }
}

// 处理用户点击事件
const handleItemClick = (type: string) => {
  switch (type) {
    case 'history':
      uni.showToast({ title: '历史记录', icon: 'none' })
      break
    case 'favorite':
      uni.showToast({ title: '我的收藏', icon: 'none' })
      break
    case 'download':
      uni.showToast({ title: '下载记录', icon: 'none' })
      break
    case 'share':
      uni.share({
        provider: 'weixin',
        scene: 'WXSceneSession',
        type: 0,
        title: '视频去水印工具',
        success: function () {
          console.log('success')
        },
        fail: function () {
          console.log('fail')
        },
      })
      break
    default:
      break
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
    default:
      break
  }
}
</script>

<style lang="scss" scoped>
:deep(.wd-button) {
  border-radius: 0.5rem;
}
</style>
