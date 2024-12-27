<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '关于我们',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-100">
    <!-- 主要内容区域 -->
    <view class="px-4 py-6 space-y-6">
      <!-- Logo和版本信息 -->
      <view class="bg-white rounded-xl p-6 shadow-sm text-center">
        <image src="/static/logo.svg" class="w-24 h-24 mx-auto mb-4" />
        <text class="text-xl font-medium text-gray-900 block">视频去水印工具</text>
        <text class="text-sm text-gray-500 mt-2 block">Version 1.0.0</text>
      </view>

      <!-- 功能特点 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-star text-blue-500"></view>
          </view>
          <text class="font-medium text-gray-900">功能特点</text>
        </view>
        <view class="space-y-3">
          <view class="flex items-start space-x-3">
            <view class="i-carbon-checkmark text-green-500 flex-shrink-0 mt-1"></view>
            <text class="text-gray-600 text-sm">支持多个主流短视频平台无水印下载</text>
          </view>
          <view class="flex items-start space-x-3">
            <view class="i-carbon-checkmark text-green-500 flex-shrink-0 mt-1"></view>
            <text class="text-gray-600 text-sm">简单易用，一键解析下载</text>
          </view>
          <view class="flex items-start space-x-3">
            <view class="i-carbon-checkmark text-green-500 flex-shrink-0 mt-1"></view>
            <text class="text-gray-600 text-sm">完全免费，无需注册</text>
          </view>
        </view>
      </view>

      <!-- 联系我们 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-email text-blue-500"></view>
          </view>
          <text class="font-medium text-gray-900">联系我们</text>
        </view>
        <view class="space-y-3">
          <view class="flex items-center space-x-3">
            <view class="i-carbon-email text-gray-400"></view>
            <text class="text-gray-600 text-sm">反馈邮箱：support@example.com</text>
          </view>
          <view class="flex items-center space-x-3">
            <view class="i-carbon-time text-gray-400"></view>
            <text class="text-gray-600 text-sm">工作时间：周一至周五 9:00-18:00</text>
          </view>
        </view>
      </view>

      <!-- 免责声明 -->
      <view class="bg-white rounded-xl p-6 shadow-sm">
        <view class="flex items-center space-x-3 mb-4">
          <view class="w-8 h-8 bg-amber-100 rounded-lg flex items-center justify-center">
            <view class="i-carbon-warning text-amber-500"></view>
          </view>
          <text class="font-medium text-gray-900">免责声明</text>
        </view>
        <view class="text-sm text-gray-600 space-y-2">
          <text class="block">
            本工具仅供学习交流使用，请勿用于任何商业用途。用户在使用过程中因违反相关法律法规或侵犯第三方权益所造成的损失，由用户自行承担。
          </text>
        </view>
      </view>

      <!-- 版权信息 -->
      <view class="text-center text-gray-400 text-xs py-4">
        <text>Copyright © 2024 All Rights Reserved</text>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
// 可以添加需要的逻辑代码
</script>

<style lang="scss" scoped>
// 可以添加需要的样式
</style>
