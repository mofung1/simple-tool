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
  <view class="min-h-screen bg-gray-100">
    <!-- 顶部通知栏 -->
    <view class="bg-amber-50 border-l-4 border-amber-400 p-4 mb-6">
      <wd-notice-bar
        text="本工具仅供学习交流使用，请勿用于商业用途。"
        prefix="warn-bold"
        custom-class="text-amber-700"
      />
    </view>

    <!-- 主要内容区域 -->
    <view class="px-4 space-y-6 pb-8">
      <!-- 功能介绍卡片 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-video text-blue-500 text-xl"></view>
          </view>
          <text class="text-lg font-medium">短视频去水印</text>
        </view>

        <text class="text-gray-600 text-sm mb-4 block">
          支持抖音、快手、小红书等平台的视频无水印下载
        </text>

        <view class="space-y-2 text-sm text-gray-600">
          <view class="flex items-center space-x-2">
            <view class="w-5 h-5 bg-blue-50 rounded-full flex items-center justify-center">
              <text class="text-blue-500 font-medium">1</text>
            </view>
            <text>打开视频平台，点击分享按钮</text>
          </view>
          <view class="flex items-center space-x-2">
            <view class="w-5 h-5 bg-blue-50 rounded-full flex items-center justify-center">
              <text class="text-blue-500 font-medium">2</text>
            </view>
            <text>复制分享链接到剪切板</text>
          </view>
          <view class="flex items-center space-x-2">
            <view class="w-5 h-5 bg-blue-50 rounded-full flex items-center justify-center">
              <text class="text-blue-500 font-medium">3</text>
            </view>
            <text>点击上方粘贴按钮，开始解析</text>
          </view>
        </view>
      </view>

      <!-- 输入框卡片 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="mb-4">
          <view
            class="flex items-center justify-between bg-blue-50 rounded-lg p-3 mb-2"
            hover-class="bg-blue-100"
            @click="pasteFromClipboardDirect"
          >
            <view class="flex items-center space-x-2">
              <text class="i-carbon-paste text-blue-500"></text>
              <text class="text-blue-600 text-sm">点击从剪切板粘贴</text>
            </view>
            <text class="i-carbon-chevron-right text-blue-400"></text>
          </view>
          <wd-textarea
            type="textarea"
            v-model="content"
            placeholder="在此输入或粘贴分享链接"
            clearable
            prop="content"
            custom-class="bg-gray-50 rounded-lg"
          />
        </view>

        <view class="flex space-x-4">
          <wd-button
            block
            @click="clearContent"
            custom-class="!bg-gray-100 !text-gray-700 !border-0 flex-1"
          >
            清空内容
          </wd-button>
          <wd-button
            block
            @click="pasteFromClipboard()"
            custom-class="!bg-gradient-to-r from-blue-500 to-blue-600 !border-0 flex-1"
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
</script>

<style lang="scss" scoped>
/* 添加必要的样式 */
</style>
