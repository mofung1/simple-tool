<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '解析成功',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#517CF0',
  },
}
</route>

<template>
  <view
    class="bg-white overflow-hidden pt-2 px-4"
    :style="{ marginTop: safeAreaInsets?.top + 'px' }"
  >
    <view class="text-center text-3xl mt-2 mb-8">
      <view class="flex flex-col items-center">
        <image 
          :src="data?.author?.avatar" 
          class="w-20 h-20 rounded-full mb-4"
        />
        <text class="text-lg font-bold mb-2">{{data?.author?.name}}</text>
        <text class="text-gray-600 mb-6">{{data?.title}}</text>
        
        <view class="w-full">
          <video
            :src="data?.video_url"
            :poster="data?.cover_url"
            class="w-full rounded-lg mb-4"
            controls
          />
        </view>
        <view class="w-full">
          <wd-button 
          block 
          @click="downloadVideo"
        >
          保存到相册
        </wd-button></view>  
        
      </view>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { onLoad } from "@dcloudio/uni-app"
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

const data = ref<any>()

onLoad((options)=>{
  data.value = JSON.parse(decodeURIComponent(options.data))
  console.log(data.value)
})
const downloadVideo = async () => {
  try {
    // 显示加载提示
    uni.showLoading({
      title: '正在下载...'
    })
    
    // 下载视频
    const { tempFilePath } = await uni.downloadFile({
      url: data.value.video_url
    })
    
    // 保存到相册
    await uni.saveVideoToPhotosAlbum({
      filePath: tempFilePath
    })

    uni.hideLoading()
    uni.showToast({
      title: '保存成功',
      icon: 'success'
    })
  } catch (err) {
    uni.hideLoading()
    uni.showToast({
      title: '保存失败',
      icon: 'error'
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
