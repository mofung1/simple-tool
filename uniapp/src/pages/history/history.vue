<route lang="json5">
{
  style: {
    navigationBarTitleText: '历史记录',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>

<template>
  <view class="min-h-screen bg-gray-100">
    <!-- 历史记录列表 -->
    <view class="p-4 space-y-4">
      <template v-if="records.length > 0">
        <view
          v-for="item in records"
          :key="item.id"
          class="bg-white rounded-xl shadow-sm overflow-hidden"
        >
          <view class="p-4 space-y-3">
            <!-- 用户信息 -->
            <view class="flex items-center space-x-3">
              <image :src="item.avatar" class="w-10 h-10 rounded-full bg-gray-200" />
              <view class="flex-1">
                <text class="text-gray-900 font-medium">{{ item.author }}</text>
                <text class="text-gray-500 text-sm block">{{ formatDate(item.created_at) }}</text>
              </view>
            </view>
            <!-- 内容区域 -->
            <view class="space-y-2">
              <text class="text-gray-800 line-clamp-2">{{ item.title }}</text>
              <image
                v-if="item.cover_url"
                :src="item.cover_url"
                mode="aspectFill"
                class="w-full h-48 rounded-lg bg-gray-100"
              />
            </view>
          </view>
        </view>
      </template>

      <!-- 空状态 -->
      <view v-else class="flex flex-col items-center justify-center py-16">
        <view class="i-carbon-document-blank text-gray-300 text-5xl mb-4"></view>
        <text class="text-gray-500">暂无记录</text>
      </view>

      <!-- 加载更多 -->
      <view v-if="hasMore" class="py-4 flex justify-center">
        <text class="text-gray-500 text-sm">{{ loading ? '加载中...' : '上拉加载更多' }}</text>
      </view>
      <view v-else-if="records.length > 0" class="py-4 flex justify-center">
        <text class="text-gray-500 text-sm">没有更多数据了</text>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { formatDate as formatDateUtil } from '@/utils/date'
import { getHistoryList, type Record } from '@/service/history'

// 状态管理
const records = ref<Record[]>([])
const loading = ref(false)
const pageNo = ref(1)
const pageSize = ref(10)
const total = ref(0)
const hasMore = ref(true)

// 格式化日期
const formatDate = (date: string) => {
  return formatDateUtil(date, 'yyyy-MM-dd HH:mm')
}

// 获取历史记录列表
const fetchRecords = async () => {
  if (loading.value) return
  loading.value = true
  try {
    const { data } = await getHistoryList(pageNo.value, pageSize.value)
    records.value = [...records.value, ...data.list]
    total.value = data.total
    hasMore.value = records.value.length < total.value
    pageNo.value++
  } catch (error) {
    uni.showToast({
      title: '获取数据失败',
      icon: 'none',
    })
  } finally {
    loading.value = false
  }
}

// 监听页面触底
const onReachBottom = async () => {
  if (hasMore.value) {
    await fetchRecords()
  }
}

// 下拉刷新
const onPullDownRefresh = async () => {
  pageNo.value = 1
  records.value = []
  hasMore.value = true
  await fetchRecords()
  uni.stopPullDownRefresh()
}

// 页面加载
onMounted(() => {
  fetchRecords()
})

// 暴露页面事件处理函数
defineExpose({
  onReachBottom,
  onPullDownRefresh,
})
</script>

<style lang="scss" scoped></style>
