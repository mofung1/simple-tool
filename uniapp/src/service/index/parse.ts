import { http } from '@/utils/http'

/** GET 请求 */
export const parseVideoAPI = (url: string) => {
  return http.get<any>('/video/share/url/parse?url', { url })
}
