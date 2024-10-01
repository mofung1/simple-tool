<route lang="json5">
{
  style: {
    navigationBarTitleText: '通屏+下拉刷新+自定义导航栏',
    enablePullDownRefresh: false,
    backgroundColor: '#018d71', // 这个背景色要与页面的.top-section的背景图差不多，这样下拉刷新看起来才比较协调
    'mp-weixin': {
      navigationStyle: 'custom',
    },
  },
}
</route>

<template>
  <view class="scroll-view-bg bg-gray-100 flex-1 h-full">
    <view  class="top-section" :style="{ paddingTop: safeAreaInsets?.top + 'px' }">

      <wd-img class="mt-20"   round width="80px" height="80px" src="https://image.mcode.fun/avatar.png" mode="aspectFit" custom-class="profile-img" />
    </view>
    <view class="user-info mb-[0rpx]">
      <view class="flex items-center justify-between px-[50rpx] pb-[50rpx] pt-[40rpx]">
            <view class="flex items-center">
                <view class="text-white ml-[20rpx]">
                 
                  <view class="text-2xl">小明</view>
                  <view class="text-xs mt-[18rpx]">
                        账号：99999999      
                    </view>
                </view>
            </view>
        </view>
    </view>

    <view class="bg-white">
      <fly-content :line="30" />
    </view>
  </view>
</template>

<script lang="ts" setup>
import { onPullDownRefresh } from '@dcloudio/uni-app'
import useNavbarWeixin from '@/hooks/useNavbarWeixin'

const { pages, isTabbar, onScrollToLower, safeAreaInsets } = useNavbarWeixin()

// 发现原生下拉刷新效果并不好，在微信里面只有顶部导航栏下拉才生效，页面区域下拉不生效，体验不好，结合自定义下拉刷新效果很好
onPullDownRefresh(() => {
  setTimeout(function fn() {
    console.log('refresh - onPullDownRefresh')
    // 关闭动画
    uni.stopPullDownRefresh()
  }, 1000)
})

// 当前下拉刷新状态
const isTriggered = ref(false)
// 自定义下拉刷新被触发
const onRefresherRefresh = async () => {
  // 开始动画
  isTriggered.value = true
  setTimeout(function fn() {
    console.log('refresh - onRefresherRefresh')
    // 关闭动画
    isTriggered.value = false
  }, 1000)
}
</script>

<style lang="scss" scoped>
.scroll-view-bg {
  // 这个背景色要与.top-section的背景图差不多，这样下拉刷新看起来才比较协调
  // background-color: #23c09c;
  // background-image: url(https://image.mcode.fun/background.jpg);
}

// 这个区域最好要大于200rpx，效果会更好
.top-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 300rpx;
  padding: 40rpx 0;
  line-height: 2;
  color: #4D80F0;
  background-color: #4D80F0;
  // background-image: url('https://image.mcode.fun/background-team.jpg');
  background-size: cover;
}

.fly-navbar {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9;
  width: 750rpx;
  color: #000;
  background-color: transparent;
}
</style>
