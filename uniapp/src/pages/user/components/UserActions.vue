<template>
  <view class="p-4 space-y-4">
    <!-- 常用功能 -->
    <view class="bg-white rounded-xl p-4 shadow-sm">
      <view class="grid grid-cols-4 gap-4">
        <view
          v-for="item in actions"
          :key="item.type"
          class="flex flex-col items-center"
          @tap="handleAction(item.type)"
        >
          <view
            :class="['w-12 h-12 rounded-xl flex items-center justify-center mb-2', item.bgColor]"
          >
            <view :class="[item.icon, item.iconColor, 'text-xl']"></view>
          </view>
          <text class="text-sm text-gray-700">{{ item.name }}</text>
        </view>
      </view>
    </view>

    <!-- 其他功能列表 -->
    <view class="bg-white rounded-xl shadow-sm overflow-hidden">
      <view class="divide-y divide-gray-100">
        <view
          v-for="item in menuItems"
          :key="item.type"
          class="flex items-center justify-between p-4"
          @tap="handleNavigation(item.type)"
        >
          <view class="flex items-center space-x-3">
            <view :class="[item.icon, 'text-gray-400']"></view>
            <text class="text-gray-700">{{ item.name }}</text>
          </view>
          <view class="i-carbon-chevron-right text-gray-400"></view>
        </view>

        <!-- 退出登录按钮 -->
        <view
          v-if="userStore.isLoggedIn()"
          class="flex items-center justify-between p-4"
          @tap="handleLogout"
        >
          <view class="flex items-center space-x-3">
            <view class="i-carbon-logout text-gray-400"></view>
            <text class="text-gray-700">退出登录</text>
          </view>
          <view class="i-carbon-chevron-right text-gray-400"></view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

// 功能按钮配置
const actions = [
  {
    type: 'history',
    name: '历史记录',
    icon: 'i-carbon-document',
    iconColor: 'text-blue-500',
    bgColor: 'bg-blue-50',
  },
  {
    type: 'favorite',
    name: '我的收藏',
    icon: 'i-carbon-favorite',
    iconColor: 'text-purple-500',
    bgColor: 'bg-purple-50',
  },
  {
    type: 'download',
    name: '下载记录',
    icon: 'i-carbon-download',
    iconColor: 'text-green-500',
    bgColor: 'bg-green-50',
  },
  {
    type: 'share',
    name: '分享应用',
    icon: 'i-carbon-share',
    iconColor: 'text-orange-500',
    bgColor: 'bg-orange-50',
  },
]

// 菜单项配置
const menuItems = [
  {
    type: 'help',
    name: '帮助与反馈',
    icon: 'i-carbon-help',
  },
  {
    type: 'about',
    name: '关于我们',
    icon: 'i-carbon-information',
  },
]

// 处理功能按钮点击
const handleAction = (type: string) => {
  if (!userStore.isLoggedIn()) {
    uni.showToast({
      title: '请先登录',
      icon: 'none',
    })
    return
  }

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
          console.log('分享成功')
        },
        fail: function () {
          console.log('分享失败')
        },
      })
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
  }
}

// 处理退出登录
const handleLogout = () => {
  uni.showModal({
    title: '提示',
    content: '确定要退出登录吗？',
    success: (res) => {
      if (res.confirm) {
        userStore.logout()
        uni.showToast({
          title: '已退出登录',
          icon: 'success',
        })
      }
    },
  })
}
</script>
