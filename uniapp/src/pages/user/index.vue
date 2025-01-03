<template>
  <view class="min-h-screen bg-gray-100">
    <!-- 用户信息 -->
    <user-info />
    <!-- 用户操作 -->
    <user-actions />
  </view>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import { useUserStore } from '@/stores/user'
import UserInfo from './components/UserInfo.vue'
import UserActions from './components/UserActions.vue'

const userStore = useUserStore()

// 页面显示时加载存储的信息
onShow(() => {
  userStore.loadStorageInfo()
})

// 初始化时也加载一次
onMounted(() => {
  userStore.loadStorageInfo()
})
</script>

<route lang="json5">
{
  style: {
    navigationBarTitleText: '',
    navigationBarTextStyle: 'white',
    navigationBarBackgroundColor: '#3B82F6',
  },
}
</route>
