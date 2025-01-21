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
  <view class="page-container">
    <!-- 聊天内容区域 -->
    <scroll-view
      :scroll-y="true"
      class="chat-container"
      :scroll-into-view="latestMessageId"
      :scroll-with-animation="true"
      :refresher-enabled="false"
      :show-scrollbar="false"
    >
      <view class="chat-list">
        <template v-if="messages.length === 0">
          <!-- 空状态提示 -->
          <view class="flex flex-col items-center justify-center pt-32">
            <view class="i-carbon-bot text-6xl text-gray-300 mb-4"></view>
            <text class="text-gray-400 text-lg">开始与AI助手对话吧</text>
          </view>
        </template>

        <template v-else>
          <template v-for="(message, index) in messages" :key="message.id">
            <!-- 消息时间 -->
            <view v-if="shouldShowTime(message, index)" class="message-time">
              {{ formatTime(message.timestamp) }}
            </view>
            
            <!-- AI消息 -->
            <view
              v-if="message.role === 'assistant'"
              :id="message.id"
              class="message-item left-message"
            >
              <view class="avatar">
                <view class="i-carbon-bot text-white text-xl"></view>
              </view>
              <view class="message-content">
                <view class="message-bubble">
                  <text class="message-text">{{ message.content }}</text>
                </view>
              </view>
            </view>

            <!-- 用户消息 -->
            <view
              v-else
              :id="message.id"
              class="message-item right-message"
            >
            <view class="avatar">
                <view class="i-carbon-user text-white text-xl"></view>
              </view>
              <view class="message-content">
                <view class="message-bubble">
                  <text class="message-text">{{ message.content }}</text>
                </view>
              </view>
              
            </view>
          </template>
        </template>

        <!-- 加载中提示 -->
        <view v-if="loading" class="message-item left-message">
          <view class="avatar">
            <view class="i-carbon-bot text-white text-xl"></view>
          </view>
          <view class="message-content">
            <view class="message-bubble loading-bubble">
              <view class="typing-indicator">
                <view class="dot"></view>
                <view class="dot"></view>
                <view class="dot"></view>
              </view>
            </view>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 底部输入区域 -->
    <view class="input-container">
      <view class="input-wrapper">
        <input
          v-model="inputMessage"
          type="text"
          placeholder="输入您的问题..."
          :disabled="loading"
          @confirm="sendMessage"
        />
        <button
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

// 判断是否显示时间
const shouldShowTime = (message: Message, index: number) => {
  if (index === 0) return true
  const prevMessage = messages.value[index - 1]
  const timeDiff = message.timestamp - prevMessage.timestamp
  // 如果与上一条消息时间相差超过5分钟，显示时间
  return timeDiff > 5 * 60 * 1000
}

// 格式化时间
const formatTime = (timestamp: number) => {
  const date = new Date(timestamp)
  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')
  return `${hours}:${minutes}`
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
.page-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #ededed;
}

.chat-container {
  flex: 1;
  overflow: hidden;
}

.chat-list {
  padding: 12px 16px;
}

.message-time {
  text-align: center;
  font-size: 12px;
  color: #999;
  margin: 8px 0;
}

.message-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 16px;
  gap: 12px;
}

.left-message {
  padding-right: 48px;
}

.right-message {
  padding-left: 48px;
  flex-direction: row-reverse;
}

.message-content {
  flex: 0 1 auto;
  max-width: 70%;
  display: flex;
  flex-direction: column;
}

.right-message .message-content {
  align-items: flex-end;
}

.message-bubble {
  display: inline-block;
  padding: 12px 16px;
  border-radius: 4px;
  font-size: 15px;
  line-height: 1.4;
  word-break: break-all;
}

.left-message .message-bubble {
  background-color: #fff;
}

.right-message .message-bubble {
  background-color: #95EC69;
}

.message-text {
  white-space: pre-wrap;
}

.loading-bubble {
  padding: 12px 24px;
}

.typing-indicator {
  display: flex;
  align-items: center;
  gap: 4px;
}

.dot {
  width: 6px;
  height: 6px;
  background: #999;
  border-radius: 50%;
  animation: bounce 1.4s infinite ease-in-out;
}

.dot:nth-child(1) { animation-delay: -0.32s; }
.dot:nth-child(2) { animation-delay: -0.16s; }

@keyframes bounce {
  0%, 80%, 100% { transform: scale(0); }
  40% { transform: scale(1); }
}

.input-container {
  background-color: #f7f7f7;
  padding: 8px 12px;
  border-top: 1px solid #e5e5e5;
  padding-bottom: calc(8px + env(safe-area-inset-bottom));
}

.input-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #fff;
  border-radius: 4px;
  padding: 4px 8px;
}

.input-wrapper input {
  flex: 1;
  height: 36px;
  font-size: 15px;
  background: transparent;
  border: none;
}

.input-wrapper button {
  background-color: #07C160;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 6px 16px;
  font-size: 15px;
  opacity: 0.9;
}

.input-wrapper button:active {
  opacity: 1;
}

.input-wrapper button:disabled {
  background-color: #ccc;
  opacity: 0.6;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 4px;
  background: linear-gradient(135deg, #3B82F6 0%, #2563EB 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
</style>
