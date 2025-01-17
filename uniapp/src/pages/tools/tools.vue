<route lang="json5">
{
  style: {
    navigationStyle: 'default',
    navigationBarTitleText: '更多',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="page-container pb-8">
    <view class="p-4 pb-safe">
      <!-- 页面标题 -->
      <view class="mb-6">
        <text class="text-2xl font-bold text-gray-800">生活优惠</text>
        <text class="mt-1 text-sm text-gray-500 block">精选优质服务，助您生活更便捷</text>
      </view>

      <!-- 工具卡片网格 -->
      <view class="grid grid-cols-2 gap-4">
        <!-- 美团外卖 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=meituan&sid=plugin"
          hover-class="opacity-90"
        >
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-yellow-400 to-orange-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-restaurant text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-yellow-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-yellow-600 font-medium">立减优惠</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">美团外卖</text>
                <text class="text-xs text-gray-500">外卖红包天天领</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 饿了么外卖 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=ele&sid=plugin"
          hover-class="opacity-90"
        >
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-blue-400 to-indigo-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-delivery text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-blue-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-blue-600 font-medium">省30元</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">饿了么外卖</text>
                <text class="text-xs text-gray-500">最高省30元</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 打车出行 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=didi&sid=plugin"
          hover-class="opacity-90"
        >
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-orange-400 to-red-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-car text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-orange-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-orange-600 font-medium">特惠券</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">打车出行</text>
                <text class="text-xs text-gray-500">出行特惠券</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 酒店住宿 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=hotel&sid=plugin"
          hover-class="opacity-90"
        >
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-purple-400 to-fuchsia-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-hotel text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-purple-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-purple-600 font-medium">特惠住宿</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">酒店住宿</text>
                <text class="text-xs text-gray-500">特惠酒店优选</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 电商优惠 -->
        <navigator
          url="plugin://cps/dianshang?pub_id=204936&source=jd,taobao,douyin,pdd,vip&coupon=0&shareRate=0.9"
          hover-class="opacity-90"
        >
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-pink-400 to-rose-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-shopping-cart text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-pink-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-pink-600 font-medium">多平台</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">电商优惠</text>
                <text class="text-xs text-gray-500">多平台好价汇总</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 京东精选 -->
        <navigator url="plugin://cps/jingdong?pub_id=204936&eliteId=10" hover-class="opacity-90">
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-red-400 to-pink-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-shopping-cart text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-red-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-red-600 font-medium">品质精选</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">京东精选</text>
                <text class="text-xs text-gray-500">品质好物推荐</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 抖音团购 -->
        <navigator url="plugin://cps/douyin?pub_id=204936" hover-class="opacity-90">
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-teal-400 to-emerald-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-store text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-teal-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-teal-600 font-medium">本地好价</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">抖音团购</text>
                <text class="text-xs text-gray-500">本地生活好价</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 抖音商城 -->
        <navigator url="plugin://cps/tiktok?pub_id=204936" hover-class="opacity-90">
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-indigo-400 to-violet-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-shopping-bag text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-indigo-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-indigo-600 font-medium">好物精选</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">抖音商城</text>
                <text class="text-xs text-gray-500">抖音好物精选</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 美团券包 -->
        <navigator url="plugin://cps/coupon?pub_id=204936" hover-class="opacity-90">
          <view
            class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
          >
            <view class="flex flex-col space-y-3">
              <view class="flex items-start justify-between">
                <view
                  class="w-14 h-14 bg-gradient-to-br from-amber-400 to-yellow-500 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                >
                  <view class="i-carbon-ticket text-white text-2xl"></view>
                </view>
                <view
                  class="px-2.5 py-1 bg-amber-50 rounded-full transform transition-transform group-hover:scale-105"
                >
                  <text class="text-xs text-amber-600 font-medium">超值券包</text>
                </view>
              </view>
              <view class="transform transition-all group-hover:translate-x-1">
                <text class="text-base font-medium text-gray-900 block mb-1">美团券包</text>
                <text class="text-xs text-gray-500">超值券包领取</text>
              </view>
            </view>
          </view>
        </navigator>
      </view>
    </view>
  </view>
</template>

<style>
.page-container {
  min-height: 100vh;
  background: linear-gradient(
    135deg,
    rgba(239, 246, 255, 0.9) 0%,
    rgba(255, 255, 255, 0.95) 45%,
    rgba(254, 242, 242, 0.8) 100%
  );
}
</style>

<script lang="ts" setup>
// 分享配置
const shareConfig = {
  title: '消印乐-去水印',
  summary: '一个简单好用的去水印工具',
  imageUrl: '/static/logo.png',
  path: '/pages/tools/tools',
}

// 小程序分享消息
const onShareAppMessage = (options: Page.ShareAppMessageOption): Page.CustomShareContent => {
  console.log('分享给好友触发:', options)
  return {
    title: shareConfig.title,
    desc: shareConfig.summary,
    path: '/pages/tools/tools',
    imageUrl: shareConfig.imageUrl, // 可选：分享图片
  }
}

// 小程序分享到朋友圈
const onShareTimeline = (): Page.ShareTimelineContent => {
  console.log('分享到朋友圈触发')
  return {
    title: shareConfig.title,
    query: 'source=timeline',
    imageUrl: shareConfig.imageUrl, // 可选：分享图片
  }
}

// 暴露页面事件处理函数
defineExpose({
  onShareAppMessage,
  onShareTimeline,
})
</script>
