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
      <!-- 搜索框 -->
      <view class="mb-6">
        <view class="flex items-center bg-gray-100 rounded-full px-4 py-2">
          <view class="i-carbon-search text-gray-400 mr-2"></view>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索优惠服务"
            class="flex-1 bg-transparent text-sm text-gray-800 placeholder-gray-400 outline-none"
            @input="handleSearch"
          />
        </view>
      </view>

      <!-- 工具卡片网格 -->
      <view class="grid grid-cols-2 gap-4">
        <template v-for="item in filteredTools" :key="item.id">
          <navigator :url="item.url" hover-class="opacity-90">
            <view
              class="group bg-white rounded-2xl p-4 border border-gray-100 shadow-sm transition-all duration-300 hover:shadow-lg hover:border-blue-100 hover:-translate-y-1"
            >
              <view class="flex flex-col space-y-3">
                <view class="flex items-start justify-between">
                  <view
                    :class="item.iconBg"
                    class="w-14 h-14 rounded-xl flex items-center justify-center transform transition-transform group-hover:scale-110 group-hover:rotate-3"
                  >
                    <view :class="item.icon" class="text-white text-2xl"></view>
                  </view>
                  <view
                    :class="item.tagBg"
                    class="px-2.5 py-1 rounded-full transform transition-transform group-hover:scale-105"
                  >
                    <text :class="item.tagTextColor" class="text-xs font-medium">{{ item.tag }}</text>
                  </view>
                </view>
                <view class="transform transition-all group-hover:translate-x-1">
                  <text class="text-base font-medium text-gray-900 block mb-1">{{ item.title }}</text>
                  <text class="text-xs text-gray-500">{{ item.description }}</text>
                </view>
              </view>
            </view>
          </navigator>
        </template>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Tool {
  id: number
  title: string
  description: string
  url: string
  icon: string
  iconBg: string
  tag: string
  tagBg: string
  tagTextColor: string
}

// 定义工具列表数据
const toolsList: Tool[] = [
  {
    id: 1,
    title: '美团外卖',
    description: '外卖红包天天领',
    url: 'plugin://cps/shop?pub_id=204936&type=meituan&sid=plugin',
    icon: 'i-carbon-restaurant',
    iconBg: 'bg-gradient-to-br from-yellow-400 to-orange-500',
    tag: '立减优惠',
    tagBg: 'bg-yellow-50',
    tagTextColor: 'text-yellow-600',
  },
  {
    id: 2,
    title: '饿了么外卖',
    description: '最高省30元',
    url: 'plugin://cps/shop?pub_id=204936&type=ele&sid=plugin',
    icon: 'i-carbon-delivery',
    iconBg: 'bg-gradient-to-br from-blue-400 to-indigo-500',
    tag: '省30元',
    tagBg: 'bg-blue-50',
    tagTextColor: 'text-blue-600',
  },
  {
    id: 3,
    title: '打车出行',
    description: '特惠券天天领',
    url: 'plugin://cps/shop?pub_id=204936&type=didi&sid=plugin',
    icon: 'i-carbon-car',
    iconBg: 'bg-gradient-to-br from-orange-400 to-red-500',
    tag: '特惠券',
    tagBg: 'bg-orange-50',
    tagTextColor: 'text-orange-600',
  },
  {
    id: 4,
    title: '酒店住宿',
    description: '特惠住宿',
    url: 'plugin://cps/shop?pub_id=204936&type=hotel&sid=plugin',
    icon: 'i-carbon-hotel',
    iconBg: 'bg-gradient-to-br from-purple-400 to-fuchsia-500',
    tag: '特惠住宿',
    tagBg: 'bg-purple-50',
    tagTextColor: 'text-purple-600',
  },
  {
    id: 5,
    title: '电商优惠',
    description: '多平台好价汇总',
    url: 'plugin://cps/dianshang?pub_id=204936&source=jd,taobao,douyin,pdd,vip&coupon=0&shareRate=0.9',
    icon: 'i-carbon-shopping-cart',
    iconBg: 'bg-gradient-to-br from-pink-400 to-rose-500',
    tag: '多平台',
    tagBg: 'bg-pink-50',
    tagTextColor: 'text-pink-600',
  },
  {
    id: 6,
    title: '京东精选',
    description: '品质好物推荐',
    url: 'plugin://cps/jingdong?pub_id=204936&eliteId=10',
    icon: 'i-carbon-shopping-cart',
    iconBg: 'bg-gradient-to-br from-red-400 to-pink-500',
    tag: '品质精选',
    tagBg: 'bg-red-50',
    tagTextColor: 'text-red-600',
  },
  {
    id: 7,
    title: '抖音团购',
    description: '本地生活好价',
    url: 'plugin://cps/douyin?pub_id=204936',
    icon: 'i-carbon-store',
    iconBg: 'bg-gradient-to-br from-teal-400 to-emerald-500',
    tag: '本地好价',
    tagBg: 'bg-teal-50',
    tagTextColor: 'text-teal-600',
  },
  {
    id: 8,
    title: '抖音商城',
    description: '抖音好物精选',
    url: 'plugin://cps/tiktok?pub_id=204936',
    icon: 'i-carbon-shopping-bag',
    iconBg: 'bg-gradient-to-br from-indigo-400 to-violet-500',
    tag: '好物精选',
    tagBg: 'bg-indigo-50',
    tagTextColor: 'text-indigo-600',
  },
  {
    id: 9,
    title: '美团券包',
    description: '超值券包领取',
    url: 'plugin://cps/coupon?pub_id=204936',
    icon: 'i-carbon-ticket',
    iconBg: 'bg-gradient-to-br from-amber-400 to-yellow-500',
    tag: '超值券包',
    tagBg: 'bg-amber-50',
    tagTextColor: 'text-amber-600',
  },
]

// 搜索查询
const searchQuery = ref('')

// 过滤后的工具列表
const filteredTools = computed(() => {
  const query = searchQuery.value.toLowerCase().trim()
  if (!query) return toolsList

  return toolsList.filter(
    (tool) =>
      tool.title.toLowerCase().includes(query) || tool.description.toLowerCase().includes(query),
  )
})

// 处理搜索输入
const handleSearch = () => {
  // 可以在这里添加额外的搜索逻辑
}

// 分享相关方法
const onShareAppMessage = (options: any) => {
  return {
    title: '生活优惠',
    path: '/pages/tools/tools',
  }
}

const onShareTimeline = () => {
  return {
    title: '生活优惠',
  }
}

defineExpose({
  onShareAppMessage,
  onShareTimeline,
})
</script>

<style>
.page-container {
  min-height: 100vh;
  background-color: #f8fafc;
}
</style>
