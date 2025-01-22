<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: 'AI聊天',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="h-screen bg-gray-100 flex flex-col">
    <view class="flex-1 flex justify-center w-full bg-white relative">
      <!-- 聊天消息列表 -->
      <scroll-view
        class="flex-1 w-full max-w-[800px] px-3"
        scroll-y
        scroll-with-animation
        :scroll-into-view="scrollToView"
      >
        <view class="flex flex-col space-y-3 py-3">
          <!-- AI消息 -->
          <view class="flex items-start space-x-2 animate-fadeIn">
            <view
              class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center shrink-0 text-xs shadow-md"
            >
              AI
            </view>
            <view
              class="bg-white p-2.5 rounded-2xl rounded-tl-sm break-words max-w-[70%] shadow-lg"
            >
              <text class="text-sm leading-normal">你好！我是你的AI助手，有什么可以帮你的吗？</text>
            </view>
          </view>

          <!-- 用户消息 -->
          <view
            class="flex flex-row-reverse items-start space-x-reverse space-x-2 self-end animate-fadeIn"
          >
            <view
              class="w-8 h-8 rounded-full bg-gray-200 flex items-center justify-center shrink-0 text-xs shadow-md"
            >
              我
            </view>
            <view
              class="bg-blue-500 text-white p-2.5 rounded-2xl rounded-tr-sm break-words max-w-[70%] shadow-lg"
            >
              <text class="text-sm leading-normal">你好！</text>
            </view>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 底部输入区域 -->
    <view class="px-3 py-2 bg-white border-t border-gray-200 safe-area-bottom">
      <view class="flex space-x-2 items-center max-w-[800px] mx-auto">
        <input
          type="text"
          class="flex-1 h-10 bg-gray-100 rounded-full px-3 text-sm"
          placeholder="请输入消息..."
          v-model="inputMessage"
          @confirm="sendMessage"
        />
        <button
          class="w-16 h-10 bg-blue-500 text-white rounded-full flex items-center justify-center text-sm"
          @click="sendMessage"
        >
          发送
        </button>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
/**
 * AI聊天页面
 * @description 提供与AI助手的实时对话功能
 */
import { ref, nextTick } from 'vue'

// 输入框的消息内容
const inputMessage = ref('')
// 用于消息列表自动滚动
const scrollToView = ref('')

// 发送消息方法
const sendMessage = () => {
  if (!inputMessage.value.trim()) return
  // TODO: 实现发送消息逻辑
  inputMessage.value = ''
  // 滚动到最新消息
  nextTick(() => {
    // TODO: 设置scrollToView的值为最新消息的ID
  })
}
</script>

<style lang="scss" scoped>
.animate-fadeIn {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.safe-area-bottom {
  padding-bottom: constant(safe-area-inset-bottom);
  padding-bottom: env(safe-area-inset-bottom);
}
</style>
