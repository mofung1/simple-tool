import { http } from '@/utils/http'

/** GET 请求 */
export const parseVideoAPI = (url: string) => {
  return http.get<any>('/api/v1/parse/url', { url })
}
