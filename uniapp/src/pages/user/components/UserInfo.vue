<template>
  <view class="bg-gradient-to-b from-blue-500 to-blue-600 p-6">
    <!-- 已登录状态 -->
    <view v-if="userStore.isLoggedIn()" class="flex items-center space-x-4">
      <image
        :src="userStore.userInfo.avatar || defaultAvatar"
        class="w-16 h-16 bg-white rounded-full border-4 border-blue-400/30"
      />
      <view class="flex-1">
        <text class="text-white text-lg font-medium block">
          {{ userStore.userInfo.nickname || '未设置昵称' }}
        </text>
      </view>
      <view class="i-carbon-chevron-right text-white"></view>
    </view>

    <!-- 未登录状态 -->
    <view v-else class="flex items-center space-x-4" @tap="handleLogin">
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
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const defaultAvatar = '/static/images/default-avatar.png'

// 处理登录
const handleLogin = async () => {
  try {
    uni.showLoading({
      title: '登录中...',
      mask: true
    })

    // 先获取用户信息
    await userStore.getUserProfile()
    
    // 再执行登录
    await userStore.login()

    uni.hideLoading()
    uni.showToast({
      title: '登录成功',
      icon: 'success'
    })
  } catch (error: any) {
    uni.hideLoading()
    // 如果是用户取消，不显示错误提示
    if (error.errMsg?.includes('cancel')) {
      return
    }
    uni.showToast({
      title: error.message || '登录失败',
      icon: 'error'
    })
    console.error('登录失败:', error)
  }
}
</script>
