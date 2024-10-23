<route lang="json5" type="home">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '免费去水印',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#517CF0',
  },
}
</route>

<template>
  <view class="h-screen w-screen bg-slate-200 pb-1">
    <wd-notice-bar
      text="本工具仅供学习交流使用，请勿用于商业用途。"
      prefix="warn-bold"
      custom-class="space"
    />
    <view class="mt-5 px-4">
      <wd-cell-group custom-class="group" title="分享链接">
        <wd-textarea
          type="textarea"
          v-model="content"
          placeholder="直接粘贴短视频分享链接"
          clearable
          prop="content"
        />
      </wd-cell-group>
      <view class="mt-5">
        <wd-button block @click="pasteFromClipboard()">一键去水印</wd-button>
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const content = ref('')

// 从剪切板粘贴
const pasteFromClipboard = async () => {
  try {
  } catch (err) {
    console.error('错误:', err)
  }
}
</script>

<style lang="scss" scoped>
/* 添加必要的样式 */
</style>
