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
  <view class="min-h-screen bg-gray-50">
    <!-- 工具列表 -->
    <view class="px-4 py-4">
      <!-- 工具卡片网格 -->
      <view class="grid grid-cols-2 gap-4">
        <!-- 美团外卖 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=meituan&sid=plugin"
          hover-class="opacity-90"
        >
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-yellow-500/10 to-yellow-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-restaurant text-yellow-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">美团外卖</text>
                <text class="text-xs text-gray-500 block">外卖红包天天领</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 饿了么外卖 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=ele&sid=plugin"
          hover-class="opacity-90"
        >
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-blue-500/10 to-blue-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-delivery text-blue-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">饿了么外卖</text>
                <text class="text-xs text-gray-500 block">最高省30元</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 打车出行 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=didi&sid=plugin"
          hover-class="opacity-90"
        >
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-orange-500/10 to-orange-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-car text-orange-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">打车出行</text>
                <text class="text-xs text-gray-500 block">出行特惠券</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 酒店住宿 -->
        <navigator
          url="plugin://cps/shop?pub_id=204936&type=hotel&sid=plugin"
          hover-class="opacity-90"
        >
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-purple-500/10 to-purple-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-hotel text-purple-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">酒店住宿</text>
                <text class="text-xs text-gray-500 block">特惠酒店优选</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 电商优惠 -->
        <navigator
          url="plugin://cps/dianshang?pub_id=204936&source=jd,taobao,douyin,pdd,vip&coupon=0&shareRate=0.9"
          hover-class="opacity-90"
        >
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-pink-500/10 to-pink-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-shopping-bag text-pink-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">电商优惠</text>
                <text class="text-xs text-gray-500 block">多平台好价汇总</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 京东精选 -->
        <navigator url="plugin://cps/jingdong?pub_id=204936&eliteId=10" hover-class="opacity-90">
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-red-500/10 to-red-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-shopping-cart text-red-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">京东精选</text>
                <text class="text-xs text-gray-500 block">品质好物推荐</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 抖音团购 -->
        <navigator url="plugin://cps/douyin?pub_id=204936" hover-class="opacity-90">
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-teal-500/10 to-teal-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-store text-teal-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">抖音团购</text>
                <text class="text-xs text-gray-500 block">本地生活好价</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 抖音商城 -->
        <navigator url="plugin://cps/tiktok?pub_id=204936" hover-class="opacity-90">
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-indigo-500/10 to-indigo-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-shopping-bag text-indigo-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">抖音商城</text>
                <text class="text-xs text-gray-500 block">抖音好物精选</text>
              </view>
            </view>
          </view>
        </navigator>

        <!-- 美团券包 -->
        <navigator url="plugin://cps/coupon?pub_id=204936" hover-class="opacity-90">
          <view class="bg-white rounded-2xl p-4 shadow-sm border border-gray-100">
            <view class="flex flex-col items-center space-y-2">
              <view
                class="w-12 h-12 bg-gradient-to-br from-amber-500/10 to-amber-600/10 rounded-xl flex items-center justify-center"
              >
                <view class="i-carbon-ticket text-amber-500 text-2xl"></view>
              </view>
              <view class="text-center">
                <text class="text-base font-medium text-gray-900">美团券包</text>
                <text class="text-xs text-gray-500 block">超值券包领取</text>
              </view>
            </view>
          </view>
        </navigator>
      </view>

      <!-- 底部提示 -->
      <view class="mt-8 text-center text-sm text-gray-500">更多优惠持续添加中...</view>
    </view>
  </view>
</template>

<script lang="ts" setup>
// 这里可以添加需要的逻辑代码
</script>
