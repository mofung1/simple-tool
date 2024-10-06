<!-- 使用 type="home" 属性设置首页，其他页面不需要设置，默认为page；推荐使用json5，更强大，且允许注释 -->
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
      text="本工具仅供学习交流使用，请勿用于商业用途，否则后果自负。"
      prefix="warn-bold"
    />
    <view class="mt-5">
      <wd-card title="分享链接">
        <wd-textarea v-model="value" size="large" placeholder="直接粘贴短视频分享链接" clearable />
      </wd-card>
      <view class="px-4"><wd-button block>一键去水印</wd-button></view>
    </view>
  </view>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped></style>
