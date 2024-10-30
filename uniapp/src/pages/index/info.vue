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
    <view class="text-center text-3xl mt-8">
        nihao 
      <text class="text-red-500">1111</text>
    </view>
  </view>
</template>

<script lang="ts" setup>
import { defineProps } from 'vue';
import { onLoad } from "@dcloudio/uni-app"
// 获取屏幕边界到安全区域距离
const { safeAreaInsets } = uni.getSystemInfoSync()

const data = ref<any>()

onLoad((options)=>{
  console.log("options--111")
  console.log(options.data)
  console.log("options--222")
  data.value = JSON.parse(decodeURIComponent(options.data));
  console.log("页面传递信息111")
  console.log(data.value)
  console.log("页面传递信息222")
})

</script>

<style lang="scss" scoped>
.test-css {
  // mt-4=>1rem=>16px;
  margin-top: 16px;
}
</style>
