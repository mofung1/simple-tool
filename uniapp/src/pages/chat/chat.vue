<!--
 * @Description: AI助手聊天界面
 * @Author: mofung
 * @Date: 2024-01-21
-->
<template>
  <view class="flex flex-col h-screen bg-gray-100">
    <!-- 顶部导航栏 -->
    <view class="flex items-center justify-between px-4 py-3 bg-blue-500">
      <text class="text-white text-lg font-medium">AI助手</text>
      <view class="flex items-center space-x-2">
        <text class="i-carbon-overflow-menu-horizontal text-white text-xl"></text>
        <text class="i-carbon-scan text-white text-xl"></text>
      </view>
    </view>

    <!-- 聊天内容区域 -->
    <scroll-view
      class="flex-1 px-4 py-2"
      scroll-y
      :scroll-into-view="scrollToView"
      :scroll-with-animation="true"
    >
      <view v-if="!chatList.length" class="flex flex-col items-center justify-center h-full">
        <view class="i-carbon-bot text-gray-300 text-8xl mb-4"></view>
        <text class="text-gray-500">开始与AI助手对话吧</text>
      </view>

      <template v-else>
        <view v-for="(item, index) in chatList" :key="index" class="mb-4">
          <!-- AI消息 -->
          <view v-if="item.role === 'assistant'" class="flex items-start">
            <view class="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center mr-2">
              <view class="i-carbon-bot text-white"></view>
            </view>
            <view class="bg-white rounded-lg p-3 max-w-75%">
              <text class="text-gray-800">{{ item.content }}</text>
            </view>
          </view>

          <!-- 用户消息 -->
          <view v-else class="flex items-start justify-end">
            <view class="bg-green-500 rounded-lg p-3 max-w-75%">
              <text class="text-white">{{ item.content }}</text>
            </view>
            <view class="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center ml-2">
              <view class="i-carbon-user text-white"></view>
            </view>
          </view>
        </view>
      </template>
    </scroll-view>

    <!-- 底部输入区域 -->
    <view class="flex items-center space-x-2 p-4 bg-white border-t border-gray-200">
      <input
        v-model="inputMessage"
        class="flex-1 px-4 py-2 rounded-full bg-gray-100 focus:outline-none"
        placeholder="输入您的问题..."
        @confirm="sendMessage"
      />
      <button
        class="px-6 py-2 rounded-full bg-blue-500 text-white"
        :disabled="!inputMessage"
        @click="sendMessage"
      >
        发送
      </button>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'

// 聊天消息列表
const chatList = ref<Array<{ role: 'user' | 'assistant'; content: string; id: number }>>([])

// 输入框消息
const inputMessage = ref('')

// 滚动到指定消息
const scrollToView = ref('')

// 发送消息
const sendMessage = () => {
  if (!inputMessage.value) return

  // 添加用户消息
  chatList.value.push({
    role: 'user',
    content: inputMessage.value,
    id: chatList.value.length,
  })

  // 模拟AI回复
  setTimeout(() => {
    chatList.value.push({
      role: 'assistant',
      content: '这是一个示例回复。\n目前这只是一个界面演示，后续会接入真实的AI对话功能。',
      id: chatList.value.length,
    })
  }, 500)

  // 清空输入框
  inputMessage.value = ''

  // 滚动到底部
  nextTick(() => {
    const lastMessage = chatList.value[chatList.value.length - 1]
    scrollToView.value = `message-${lastMessage.id}`
  })
}
</script>

<style>
page {
  background-color: #f3f4f6;
}
</style>
