<route lang="json5" type="home">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '免费去水印',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-50">
    <!-- 顶部通知栏 -->
    <view class="bg-amber-50/80 backdrop-blur-sm p-2 mb-2 sticky top-0 z-50">
      <wd-notice-bar
        text="本工具仅供学习交流使用，请勿用于商业用途。"
        prefix="warn-bold"
        custom-class="text-amber-700 !bg-transparent text-lg"
      />
    </view>

    <!-- 主要内容区域 -->
    <view class="px-4 space-y-4 pb-8">
      <!-- 功能介绍卡片 -->
      <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
        <view class="flex items-center space-x-3 mb-4">
          <view
            class="w-12 h-12 bg-gradient-to-br from-blue-500/10 to-blue-600/10 rounded-xl flex items-center justify-center"
          >
            <view class="i-carbon-video text-blue-500 text-2xl"></view>
          </view>
          <text class="text-lg font-semibold text-gray-900">短视频去水印</text>
        </view>

        <text class="text-gray-600 text-sm mb-2 block leading-relaxed">
          抖音、快手、小红书等平台的视频无水印下载，一键解析快速便捷。
        </text>

        <!-- <view class="space-y-3 text-sm text-gray-600">
          <view class="flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-50 transition-colors">
            <view class="w-6 h-6 bg-blue-50 rounded-full flex items-center justify-center shrink-0">
              <text class="text-blue-500 font-medium">1</text>
            </view>
            <text>打开视频平台，点击分享按钮</text>
          </view>
          <view class="flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-50 transition-colors">
            <view class="w-6 h-6 bg-blue-50 rounded-full flex items-center justify-center shrink-0">
              <text class="text-blue-500 font-medium">2</text>
            </view>
            <text>复制分享链接到剪切板</text>
          </view>
          <view class="flex items-center space-x-3 p-2 rounded-lg hover:bg-gray-50 transition-colors">
            <view class="w-6 h-6 bg-blue-50 rounded-full flex items-center justify-center shrink-0">
              <text class="text-blue-500 font-medium">3</text>
            </view>
            <text>点击粘贴按钮，开始解析</text>
          </view>
        </view> -->
      </view>

      <!-- 输入框卡片 -->
      <view class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
        <view class="mb-4">
          <view
            class="flex items-center justify-between bg-gradient-to-r from-blue-50 to-blue-100/50 rounded-xl p-4 mb-3"
            hover-class="opacity-90"
            @click="pasteFromClipboardDirect"
          >
            <view class="flex items-center space-x-3">
              <text class="i-carbon-paste text-blue-500 text-xl"></text>
              <text class="text-blue-600 font-medium">点击从剪切板粘贴</text>
            </view>
            <text class="i-carbon-chevron-right text-blue-400"></text>
          </view>
          <wd-textarea
            type="textarea"
            v-model="content"
            placeholder="在此输入或粘贴分享链接"
            clearable
            prop="content"
            custom-class="bg-gray-50/50 rounded-xl border-gray-200"
          />
        </view>

        <view class="flex space-x-4">
          <wd-button
            block
            @click="clearContent"
            custom-class="!bg-gray-100 !text-gray-700 !border-0 flex-1 !rounded-xl !h-11"
          >
            清空内容
          </wd-button>
          <wd-button
            block
            @click="pasteFromClipboard()"
            custom-class="!bg-gradient-to-r from-blue-500 to-blue-600 !border-0 flex-1 !rounded-xl !h-11 !font-medium"
          >
            一键去水印
          </wd-button>
        </view>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { parseVideoAPI } from '@/service/index/parse'

const content = ref('')

// 直接从剪切板粘贴
const pasteFromClipboardDirect = () => {
  uni.getClipboardData({
    success: (res) => {
      if (res.data) {
        content.value = res.data
      } else {
        uni.showToast({
          title: '剪切板为空',
          icon: 'error',
        })
      }
    },
    fail: () => {
      uni.showToast({
        title: '粘贴失败',
        icon: 'error',
      })
    },
  })
}

// 清空内容
const clearContent = () => {
  content.value = ''
}

// 检查剪贴板内容
const checkClipboard = () => {
  uni.getClipboardData({
    success: (res) => {
      if (res.data && res.data.includes('http')) {
        uni.showModal({
          title: '发现链接',
          content: '检测到剪贴板中包含链接，是否使用该链接？',
          confirmText: '使用',
          cancelText: '取消',
          success: (modalRes) => {
            if (modalRes.confirm) {
              content.value = res.data
            }
          },
        })
      }
    },
    fail: (err) => {
      console.error('获取剪贴板内容失败:', err)
    },
  })
}

// 页面加载时检查剪贴板
onMounted(() => {
  // 延迟一下执行，确保页面完全加载
  setTimeout(() => {
    checkClipboard()
  }, 500)
})

// 处理粘贴并去水印的函数
const pasteFromClipboard = async () => {
  if (!content.value) {
    // 如果内容为空，提示用户粘贴链接
    uni.showToast({
      icon: 'none',
      title: '请粘贴短视频分享链接',
    })
    return
  }
  const res = await parseVideoAPI(content.value)

  if (!res.data.video_url) {
    // 如果没有视频链接，提示解析失败
    uni.showToast({
      icon: 'none',
      title: '解析失败，请检查链接是否有效',
    })
  } else {
    // 解析成功，跳转到结果页面
    const data = res.data
    uni.navigateTo({
      url: `/pages/index/info?data=${encodeURIComponent(JSON.stringify(data))}`,
    })
  }
}

// 分享配置
const shareConfig = {
  title: '消印乐-去水印',
  summary: '一个简单好用的去水印工具',
  imageUrl: '/static/logo.png',
  path: '/pages/index/index',
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
</script>

<style lang="scss" scoped>
/* 添加必要的样式 */
</style>
