<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '帮助与反馈',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-100">
    <!-- 主要内容区域 -->
    <view class="px-4 py-6 space-y-6">
      <!-- 常见问题 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-help text-blue-500"></view>
          </view>
          <text class="font-medium text-gray-900">常见问题</text>
        </view>

        <view class="space-y-4">
          <view
            v-for="(qa, index) in faqs"
            :key="index"
            class="border-b border-gray-100 last:border-0 pb-4 last:pb-0"
          >
            <view class="flex items-start space-x-3 mb-2">
              <view
                class="w-6 h-6 bg-blue-50 rounded-full flex items-center justify-center flex-shrink-0"
              >
                <text class="text-blue-500 text-sm">Q</text>
              </view>
              <text class="text-gray-900 font-medium">{{ qa.question }}</text>
            </view>
            <view class="flex items-start space-x-3 ml-9">
              <text class="text-gray-600 text-sm">{{ qa.answer }}</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 联系我们 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-email text-blue-500"></view>
          </view>
          <text class="font-medium text-gray-900">联系我</text>
        </view>
        <view class="space-y-3">
          <view class="flex items-center space-x-3">
            <view class="i-carbon-email text-gray-400"></view>
            <text class="text-gray-600 text-sm">邮箱：2279254178@qq.com</text>
          </view>
          <view class="flex items-center space-x-3">
            <view class="i-carbon:logo-wechat text-gray-400"></view>
            <text class="text-gray-600 text-sm">微信号：mofung1</text>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

// 常见问题列表
const faqs = [
  {
    question: '为什么视频链接无法解析？',
    answer:
      '请确保视频是公开的且链接完整有效。某些平台可能会限制视频访问，建议重新获取分享链接后再试。',
  },
  {
    question: '下载的视频没有声音怎么办？',
    answer: '这可能是由于视频源的限制导致。建议尝试重新获取视频链接，或者使用其他视频平台的链接。',
  },
  {
    question: '为什么下载速度很慢？',
    answer:
      '下载速度受网络环境和视频大小影响。建议使用WiFi网络，并确保网络连接稳定。如果视频文件较大，可能需要等待较长时间。',
  },
  {
    question: '支持哪些视频平台？',
    answer:
      '目前支持抖音、快手、小红书、微视、西瓜视频、B站等主流短视频平台。我们会持续增加对其他平台的支持。',
  },
]

// 表单数据
const feedback = ref('')
const contact = ref('')
const contactType = ref('')

// 提交反馈
const submitFeedback = () => {
  if (!feedback.value.trim()) {
    uni.showToast({
      title: '请输入反馈内容',
      icon: 'none',
    })
    return
  }

  // TODO: 实现反馈提交逻辑
  uni.showToast({
    title: '反馈提交成功',
    icon: 'success',
  })

  // 清空表单
  feedback.value = ''
  contact.value = ''
  contactType.value = ''
}
</script>

<style lang="scss" scoped>
// 可以添加需要的样式
</style>
