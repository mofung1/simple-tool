<template>
  <view class="user-avatar">
    <button v-if="!userInfo.avatar" class="avatar-btn" @tap="handleGetUserProfile">
      点击获取头像
    </button>
    <image v-else class="avatar" :src="userInfo.avatar" mode="aspectFill" />
    <text class="nickname">{{ userInfo.nickname || '未设置昵称' }}</text>
  </view>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useUserStore } from '@/stores/user'
import { getUserProfile } from '@/utils/auth'

const userStore = useUserStore()
const userInfo = computed(() => userStore.userInfo)

const handleGetUserProfile = async () => {
  try {
    await getUserProfile()
    uni.showToast({
      title: '获取成功',
      icon: 'success'
    })
  } catch (error) {
    uni.showToast({
      title: error.message || '获取失败',
      icon: 'none'
    })
  }
}
</script>

<style lang="scss">
.user-avatar {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20rpx;

  .avatar-btn {
    width: 128rpx;
    height: 128rpx;
    border-radius: 50%;
    background: #f5f5f5;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 24rpx;
    color: #666;
    margin-bottom: 16rpx;
  }

  .avatar {
    width: 128rpx;
    height: 128rpx;
    border-radius: 50%;
    margin-bottom: 16rpx;
  }

  .nickname {
    font-size: 28rpx;
    color: #333;
  }
}
</style>
