<route lang="json5">
{
  style: {
    navigationBarTitleText: '通屏+下拉刷新+自定义导航栏',
    enablePullDownRefresh: false,
    backgroundColor: '#23c09c', // 这个背景色要与页面的.top-section的背景图差不多，这样下拉刷新看起来才比较协调
    'app-plus': {
      titleNView: {
        type: 'transparent',
      },
    },
    'mp-weixin': {
      navigationStyle: 'custom',
    },
  },
}
</route>

<template>
  <view class="scroll-view-bg flex-1 h-full">
    <view class="top-section" :style="{ paddingTop: safeAreaInsets?.top + 'px' }"></view>
    <!-- <view class="p-2 leading-6 bg-white"> -->
    <wd-card title="经营分析">
      一般的，检举内容由承办的党的委员会或纪律检查委员会将处理意见或复议、复查结论同申诉人见面，听取其意见。复议、复查的结论和决定，应交给申诉人一份。
      <template #footer>
        <wd-button size="small" plain>查看详情</wd-button>
      </template>
    </wd-card>
    <!-- </view> -->
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
  background-image: url(https://image.mcode.fun/background.jpg);
}

// 这个区域最好要大于200rpx，效果会更好
.top-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 200rpx;
  padding: 40rpx 0;
  line-height: 2;
  color: #fff;
  background-image: url('https://image.mcode.fun/background-team.jpg');
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
