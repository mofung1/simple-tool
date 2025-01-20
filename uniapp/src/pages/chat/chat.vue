<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: 'AI助手',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-50">
    <!-- 聊天内容区域 -->
    <scroll-view
      :scroll-y="true"
      class="chat-container px-4 pb-32"
      :scroll-into-view="latestMessageId"
    >
      <template v-if="messages.length === 0">
        <!-- 空状态提示 -->
        <view class="flex flex-col items-center justify-center pt-32">
          <view class="i-carbon-bot text-6xl text-gray-300 mb-4"></view>
          <text class="text-gray-400 text-lg">开始与AI助手对话吧</text>
        </view>
      </template>

      <template v-else>
        <template v-for="(message, index) in messages" :key="message.id">
          <!-- AI消息 -->
          <view
            v-if="message.role === 'assistant'"
            :id="message.id"
            class="flex items-start space-x-3 mb-6"
          >
            <view
              class="w-10 h-10 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl flex items-center justify-center"
            >
              <view class="i-carbon-bot text-white text-xl"></view>
            </view>
            <view class="flex-1">
              <view
                class="bg-white rounded-2xl rounded-tl-none px-4 py-3 shadow-sm border border-gray-100"
              >
                <text class="text-gray-800 leading-relaxed whitespace-pre-wrap">
                  {{ message.content }}
                </text>
              </view>
            </view>
          </view>

          <!-- 用户消息 -->
          <view v-else :id="message.id" class="flex items-start space-x-3 mb-6 flex-row-reverse">
            <view
              class="w-10 h-10 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl flex items-center justify-center"
            >
              <view class="i-carbon-user text-white text-xl"></view>
            </view>
            <view class="flex-1">
              <view class="bg-blue-500 rounded-2xl rounded-tr-none px-4 py-3 shadow-sm">
                <text class="text-white leading-relaxed whitespace-pre-wrap">
                  {{ message.content }}
                </text>
              </view>
            </view>
          </view>
        </template>
      </template>

      <!-- 加载中提示 -->
      <view v-if="loading" class="flex items-center justify-center py-4">
        <wd-loading size="36px" color="#3B82F6" />
      </view>
    </scroll-view>

    <!-- 底部输入区域 -->
    <view class="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-100 px-4 py-3">
      <view class="flex items-center space-x-3">
        <input
          v-model="inputMessage"
          type="text"
          placeholder="输入您的问题..."
          class="flex-1 bg-gray-100 rounded-xl px-4 py-2.5 text-base text-gray-800 placeholder-gray-400"
          :disabled="loading"
          @confirm="sendMessage"
        />
        <button
          class="bg-blue-500 hover:bg-blue-600 text-white rounded-xl px-6 py-2.5 text-base font-medium disabled:opacity-50"
          :disabled="!inputMessage.trim() || loading"
          @tap="sendMessage"
        >
          发送
        </button>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 定义消息类型
interface Message {
  id: string
  role: 'user' | 'assistant'
  content: string
  timestamp: number
}

// 状态变量
const messages = ref<Message[]>([])
const inputMessage = ref('')
const loading = ref(false)
const latestMessageId = ref('')

// 生成唯一ID
const generateId = () => {
  return 'msg_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9)
}

// 发送消息
const sendMessage = async () => {
  const message = inputMessage.value.trim()
  if (!message || loading.value) return

  // 添加用户消息
  const userMessage: Message = {
    id: generateId(),
    role: 'user',
    content: message,
    timestamp: Date.now(),
  }
  messages.value.push(userMessage)
  latestMessageId.value = userMessage.id
  inputMessage.value = ''

  // 开始加载
  loading.value = true

  try {
    // 模拟AI回复
    const aiMessage: Message = {
      id: generateId(),
      role: 'assistant',
      content: '这是一个示例回复。目前这只是一个界面演示，后续会接入真实的AI对话功能。',
      timestamp: Date.now(),
    }

    // 延迟一下以模拟网络请求
    await new Promise((resolve) => setTimeout(resolve, 1000))

    messages.value.push(aiMessage)
    latestMessageId.value = aiMessage.id
  } catch (error) {
    console.error('Error:', error)
    uni.showToast({
      title: '发送失败，请重试',
      icon: 'none',
    })
  } finally {
    loading.value = false
  }
}
</script>

<style>
.chat-container {
  height: calc(100vh - 140px);
}
</style>
