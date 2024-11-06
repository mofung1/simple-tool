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
  import { parseVideoAPI } from '@/service/index/parse'
  
  const content = ref('')
  
  // 处理粘贴并去水印的函数
  const pasteFromClipboard = async () => {
    if (!content.value) {
      // 如果内容为空，提示用户粘贴链接
      uni.showToast({
        icon: 'none',
        title: '请粘贴短视频分享链接',
      })
      return
    }
  
    try {
      uni.showLoading({
        title: '解析中...'
      })
  
      const res = await parseVideoAPI(content.value)
  
      if (!res.data.video_url) {
        // 如果没有视频链接，提示解析失败
        uni.showToast({
          icon: 'none',
          title: '解析失败，请检查链接是否有效',
        })
      } else {
        // 解析成功，跳转到结果页面
        let data = res.data
        uni.navigateTo({
          url: `/pages/index/info?data=${encodeURIComponent(JSON.stringify(data))}`
        })
      }
    } catch (err) {
      // 捕获错误并显示错误提示
      uni.showToast({
        icon: 'none',
        title: err instanceof Error ? err.message : '发生了未知错误',
      })
    } finally {
      // 无论成功或失败都要隐藏加载中动画
      uni.hideLoading()
    }
  }
  </script>
  
  <style lang="scss" scoped>
  /* 添加必要的样式 */
  </style>
  