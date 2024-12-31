<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '解析成功',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view
    class="min-h-screen bg-gray-100 px-4 py-6"
    :style="{ paddingTop: safeAreaInsets?.top + 'px' }"
  >
    <!-- 作者信息卡片 -->
    <view class="bg-white rounded-xl p-6 shadow-sm mb-6">
      <view class="flex items-center space-x-4">
        <image :src="data?.author?.avatar" class="w-16 h-16 rounded-full border-4 border-gray-50" />
        <view>
          <text class="font-medium text-gray-900 block mb-1">{{ data?.author?.name }}</text>
          <text class="text-sm text-gray-500">创作者</text>
        </view>
      </view>
    </view>

    <!-- 视频内容卡片 -->
    <view class="bg-white rounded-xl p-6 shadow-sm mb-6">
      <!-- 视频标题 -->
      <view class="mb-4">
        <text class="text-sm text-gray-500">视频标题</text>
        <text class="block mt-1 text-gray-900">{{ data?.title }}</text>
      </view>

      <!-- 视频播放器 -->
      <view class="relative rounded-lg overflow-hidden bg-black aspect-video mb-6">
        <video
          :src="data?.video_url"
          :poster="data?.cover_url"
          class="w-full h-full object-contain"
          controls
        />
      </view>

      <!-- 下载按钮 -->
      <wd-button
        block
        @click="downloadVideo"
        custom-class="!bg-gradient-to-r from-blue-500 to-blue-600 !border-0"
      >
        保存到相册
      </wd-button>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { onLoad } from '@dcloudio/uni-app'
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

const data = ref<any>()

onLoad((options) => {
  data.value = JSON.parse(decodeURIComponent(options.data))
  console.log(data.value)
})
const downloadVideo = async () => {
  try {
    // 显示加载提示
    uni.showLoading({
      title: '正在下载...',
    })

    // 下载视频
    const { tempFilePath } = await uni.downloadFile({
      url: data.value.video_url,
    })

    // 保存到相册
    await uni.saveVideoToPhotosAlbum({
      filePath: tempFilePath,
    })

    uni.hideLoading()
    uni.showToast({
      title: '保存成功',
      icon: 'success',
    })
  } catch (err) {
    uni.hideLoading()
    uni.showToast({
      title: '保存失败',
      icon: 'error',
    })
  }
}
</script>

<style lang="scss" scoped>
.test-css {
  // mt-4=>1rem=>16px;
  margin-top: 16px;
}
</style>
