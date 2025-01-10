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
    <view
      class="bg-gradient-to-br from-blue-500 to-blue-600 p-6 relative overflow-hidden login-card"
    >
      <!-- 背景装饰 -->
      <view class="absolute -right-12 -top-12 w-48 h-48 rounded-full bg-blue-400/20"></view>
      <view class="absolute -left-12 -bottom-12 w-36 h-36 rounded-full bg-blue-400/10"></view>

      <view v-if="userStore.isLogined" class="flex items-center space-x-4 relative">
        <view class="relative">
          <image
            :src="userStore.userInfo.avatar"
            class="w-20 h-20 bg-white rounded-full border-4 border-white/30 shadow-lg"
          />
          <view class="absolute bottom-0 right-0 w-6 h-6 rounded-full border-2 border-white"></view>
        </view>
        <view class="flex-1">
          <text class="text-white text-xl font-semibold block mb-1">
            {{ userStore.userInfo.nickname }}
          </text>
          <text class="text-blue-100 text-sm opacity-80">ID: {{ userStore.userInfo.sn }}</text>
        </view>
        <!-- <view class="i-carbon-chevron-right text-white/70 text-xl"></view> -->
      </view>

      <view
        v-else
        :plain="true"
        class="flex items-center space-x-4 w-full border-none p-0 relative !m-0 !after:border-none"
        style="background-color: transparent !important"
      >
        <view class="flex-1">
          <view class="flex flex-col items-start">
            <button
              class="flex items-center space-x-2 mt-10 bg-white hover:bg-white/30 px-6 py-2 rounded-lg text-white text-sm transition-colors"
              :class="{ 'opacity-60': loading }"
              :disabled="loading"
              @tap="handleGetUserProfile"
            >
              <text class="text-blue">{{ loading ? '登录中...' : '点击登录' }}</text>
              <view class="i-carbon-login text-blue"></view>
            </button>
          </view>
        </view>
      </view>
    </view>

    <!-- 功能区域 -->
    <view class="p-4 space-y-4">
      <!-- 其他功能列表 -->
      <view class="bg-white rounded-xl shadow-sm overflow-hidden">
        <view class="divide-y divide-gray-100">
          <!-- <view class="flex items-center justify-between p-4" @tap="handleNavigation('history')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-document text-blue-500 text-xl"></view>
              <text class="text-gray-700">历史记录</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view> -->
          <view class="flex items-center justify-between p-4" @tap="handleNavigation('help')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-help text-purple-500 text-xl"></view>
              <text class="text-gray-700">使用帮助</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view>
          <view class="flex items-center justify-between p-4" @tap="handleNavigation('about')">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-information text-green-500 text-xl"></view>
              <text class="text-gray-700">关于我们</text>
            </view>
            <view class="i-carbon-chevron-right text-gray-400"></view>
          </view>
          <view class="flex items-center justify-between p-4" @tap="handleShare">
            <view class="flex items-center space-x-3">
              <view class="i-carbon-share text-orange-500 text-xl"></view>
              <text class="text-gray-700">分享应用</text>
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
  title: '简单工具',
  summary: '一个简单好用的视频下载工具',
  imageUrl: '/static/logo.png',
  path: '/pages/index/index'
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
        icon: 'success'
      })
    },
    fail: function () {
      uni.showToast({
        title: '分享失败',
        icon: 'error'
      })
    }
  })
  // #endif

  // #ifdef H5
  if (navigator.share) {
    navigator.share({
      title: shareConfig.title,
      text: shareConfig.summary,
      url: window.location.href
    }).catch(() => {
      uni.showToast({
        title: '分享失败',
        icon: 'error'
      })
    })
  } else {
    uni.setClipboardData({
      data: window.location.href,
      success: () => {
        uni.showToast({
          title: '链接已复制',
          icon: 'success'
        })
      }
    })
  }
  // #endif

  // #ifdef MP-WEIXIN
  uni.showToast({
    title: '点击右上角分享',
    icon: 'none'
  })
  // #endif
}

// 小程序分享消息
const onShareAppMessage = () => {
  return {
    title: shareConfig.title,
    path: shareConfig.path,
    imageUrl: shareConfig.imageUrl
  }
}

// 小程序分享到朋友圈
const onShareTimeline = () => {
  return {
    title: shareConfig.title,
    path: shareConfig.path,
    imageUrl: shareConfig.imageUrl
  }
}

// 暴露页面事件处理函数
defineExpose({
  onShareAppMessage,
  onShareTimeline
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
.login-card {
  height: 100px;
}
</style>
