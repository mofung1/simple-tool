<route lang="json5">
{
  style: {
    navigationBarTitleText: '',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
    navigationStyle: 'custom',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-50">
    <!-- 用户信息卡片 -->
    <view class="bg-gradient-to-br from-blue-500 to-blue-600 relative overflow-hidden">
      <!-- 顶部安全区域 -->
      <view class="h-12" :style="{ paddingTop: safeAreaInsets?.top + 'px' }"></view>

      <!-- 内容区域 -->
      <view class="px-8 pb-12 pt-4">
        <!-- 背景装饰 -->
        <view
          class="absolute -right-16 -top-16 w-56 h-56 rounded-full bg-blue-400/20 animate-pulse"
        ></view>
        <view
          class="absolute -left-16 -bottom-16 w-48 h-48 rounded-full bg-blue-400/10 animate-pulse-slow"
        ></view>

        <view v-if="userStore.isLogined" class="flex items-center space-x-5 relative">
          <view class="relative group">
            <image
              :src="userStore.userInfo.avatar"
              class="w-20 h-20 bg-white rounded-full border-4 border-white/40 shadow-lg transition-transform hover:scale-105"
            />
            <view
              class="absolute bottom-0 right-0 w-5 h-5 rounded-full border-2 border-white bg-green-400"
            ></view>
          </view>
          <view class="flex-1">
            <text class="text-white text-xl font-bold block mb-2">
              {{ userStore.userInfo.nickname }}
            </text>
            <view class="flex items-center space-x-2">
              <text class="text-blue-50 text-sm opacity-90">ID: {{ userStore.userInfo.sn }}</text>
              <view class="w-1.5 h-1.5 rounded-full bg-blue-50 opacity-50"></view>
            </view>
          </view>
        </view>

        <view v-else class="flex items-center space-x-4 w-full relative py-4">
          <view class="flex-1">
            <view class="flex flex-col items-start">
              <button
                class="group flex items-center space-x-3 mt-4 bg-white hover:bg-white/90 px-8 py-3 rounded-2xl text-base transition-all duration-300 shadow-lg"
                :class="{ 'opacity-60': loading }"
                :disabled="loading"
                @tap="handleGetUserProfile"
              >
                <text
                  class="text-blue-500 font-medium group-hover:translate-x-1 transition-transform"
                >
                  {{ loading ? '登录中...' : '立即登录' }}
                </text>
                <view
                  class="i-carbon-login text-blue-500 text-lg group-hover:translate-x-1 transition-transform"
                ></view>
              </button>
            </view>
          </view>
        </view>
      </view>
    </view>

    <!-- 功能区域 -->
    <view class="p-4 space-y-4">
      <!-- 其他功能列表 -->
      <view class="bg-white rounded-2xl shadow-sm overflow-hidden">
        <view class="divide-y divide-gray-50">
          <view
            v-for="(item, index) in [
              { icon: 'i-carbon-help', color: 'text-purple-500', text: '使用帮助', type: 'help' },
              {
                icon: 'i-carbon-information',
                color: 'text-green-500',
                text: '关于我们',
                type: 'about',
              },
              { icon: 'i-carbon-share', color: 'text-orange-500', text: '分享应用', type: 'share' },
            ]"
            :key="index"
            class="group flex items-center justify-between p-4 hover:bg-gray-50 transition-colors cursor-pointer"
            @tap="item.type === 'share' ? handleShare() : handleNavigation(item.type)"
          >
            <view class="flex items-center space-x-3">
              <view :class="[item.icon, item.color, 'text-xl']"></view>
              <text class="text-gray-700">{{ item.text }}</text>
            </view>
            <view
              class="i-carbon-chevron-right text-gray-400 group-hover:translate-x-1 transition-transform"
            ></view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { ref, onMounted } from 'vue'
import { checkLogin } from '@/utils/auth'

const userStore = useUserStore()

// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

const loading = ref(false)

const handleGetUserProfile = () => {
  loading.value = true
  uni.getUserProfile({
    desc: '用于完善用户资料',
    lang: 'zh_CN',
    success: async (profileRes) => {
      try {
        // 获取登录凭证
        const { code } = await uni.login({
          provider: 'weixin',
        })

        console.log('登录凭证：', code)
        console.log('用户信息：', profileRes.userInfo)

        // 调用后端登录接口
        uni.showLoading({ title: '登录中...' })
        await userStore.wxLogin({
          code,
          userInfo: profileRes.userInfo,
        })
        uni.hideLoading()
      } catch (error: any) {
        uni.showToast({
          title: '登录失败',
          icon: 'error',
        })
      } finally {
        loading.value = false
      }
    },
    fail: (error) => {
      loading.value = false
      console.error('获取用户信息失败：', error)
      if (error.errMsg?.includes('getUserProfile:fail auth deny')) {
        uni.showToast({
          title: '需要您的授权才能继续',
          icon: 'none',
        })
      } else {
        uni.showToast({
          title: '登录失败',
          icon: 'error',
        })
      }
    },
  })
}

// 处理页面导航
const handleNavigation = (type: string) => {
  switch (type) {
    case 'help':
      uni.navigateTo({
        url: '/pages/help/help',
      })
      break
    case 'about':
      uni.navigateTo({
        url: '/pages/about/about',
      })
      break
    case 'history':
      if (checkLogin()) {
        uni.navigateTo({
          url: '/pages/history/history',
        })
      }
      break
  }
}

// 分享配置
const shareConfig = {
  title: '消印乐-去水印',
  summary: '一个简单好用的去水印工具',
  imageUrl: '/static/logo.png',
  path: '/pages/index/index',
}

// 处理分享
const handleShare = () => {
  // #ifdef APP-PLUS
  uni.share({
    provider: 'weixin',
    scene: 'WXSceneSession',
    type: 0,
    ...shareConfig,
    success: function () {
      uni.showToast({
        title: '分享成功',
        icon: 'success',
      })
    },
    fail: function () {
      uni.showToast({
        title: '分享失败',
        icon: 'error',
      })
    },
  })
  // #endif

  // #ifdef H5
  if (navigator.share) {
    navigator
      .share({
        title: shareConfig.title,
        text: shareConfig.summary,
        url: window.location.href,
      })
      .catch(() => {
        uni.showToast({
          title: '分享失败',
          icon: 'error',
        })
      })
  } else {
    uni.setClipboardData({
      data: window.location.href,
      success: () => {
        uni.showToast({
          title: '链接已复制',
          icon: 'success',
        })
      },
    })
  }
  // #endif

  // #ifdef MP-WEIXIN
  uni.showToast({
    title: '点击右上角分享',
    icon: 'none',
  })
  // #endif
}

// 小程序分享消息
const onShareAppMessage = (options: Page.ShareAppMessageOption): Page.CustomShareContent => {
  console.log('分享给好友触发:', options)
  return {
    title: shareConfig.title,
    desc: shareConfig.summary,
    path: '/pages/index/index',
    imageUrl: shareConfig.imageUrl, // 可选：分享图片
  }
}

// 小程序分享到朋友圈
const onShareTimeline = (): Page.ShareTimelineContent => {
  console.log('分享到朋友圈触发')
  return {
    title: shareConfig.title,
    query: 'source=timeline',
    imageUrl: shareConfig.imageUrl, // 可选：分享图片
  }
}

// 暴露页面事件处理函数
defineExpose({
  onShareAppMessage,
  onShareTimeline,
})

// 页面加载
onMounted(() => {
  // #ifdef MP-WEIXIN
  // #endif
})
</script>

<style lang="scss" scoped>
:deep(.uni-button) {
  &::after {
    border: none;
  }
}
:deep(.wd-button) {
  border-radius: 0.5rem;
}
</style>
