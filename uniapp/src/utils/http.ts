import { CustomRequestOptions } from '@/interceptors/request'
import { useUserStore } from '@/stores/user'

export const http = <T>(options: CustomRequestOptions) => {
  // 1. 返回 Promise 对象
  return new Promise<IResData<T>>((resolve, reject) => {
    uni.request({
      ...options,
      dataType: 'json',
      // #ifndef MP-WEIXIN
      responseType: 'json',
      // #endif
      // 响应成功
      success(res) {
        console.log(res)
        if (res.statusCode >= 200 && res.statusCode < 300) {
          resolve(res.data as IResData<T>)
        } else {
          if (res.statusCode === 401) {
            const userStore = useUserStore()
            userStore.clearUserInfo()
            uni.navigateTo({ url: '/pages/index/index' })
          } else {
            // 统一的错误处理
            let errorMsg = '请求错误'
            if ((res.data as IResData<T>).code !== 200) {
              errorMsg = (res.data as IResData<T>).msg
            }
            !options.hideErrorToast &&
              uni.showToast({
                icon: 'none',
                title: errorMsg,
              })
          }
          reject(res)
        }
      },
      // 响应失败
      fail(err) {
        uni.showToast({
          icon: 'none',
          title: '网络错误，请检查网络连接',
        })
        reject(err)
      },
    })
  })
}

/**
 * GET 请求
 * @param url 后台地址
 * @param query 请求query参数
 * @returns
 */
export const httpGet = <T>(url: string, query?: Record<string, any>) => {
  return http<T>({
    url,
    method: 'GET',
    query,
  })
}

/**
 * POST 请求
 * @param url 后台地址
 * @param data 请求body参数
 * @param query 请求query参数
 * @returns
 */
export const httpPost = <T>(
  url: string,
  data?: Record<string, any>,
  query?: Record<string, any>,
) => {
  return http<T>({
    url,
    method: 'POST',
    data,
    query,
  })
}

http.get = httpGet
http.post = httpPost
