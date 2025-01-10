/**
 * 历史记录相关接口
 */

import { http } from '@/utils/http'

export interface Record {
  id: number
  created_at: string
  updated_at: string
  deleted_at: string | null
  user_id: number
  author: string
  avatar: string
  title: string
  cover_url: string
  music_url: string
  video_url: string
}

export interface ListResponse {
  code: number
  msg: string
  data: {
    total: number
    page_no: number
    page_size: number
    list: Record[]
  }
}

/**
 * 获取历史记录列表
 * @param pageNo - 页码
 * @param pageSize - 每页数量
 */
export const getHistoryList = (pageNo: number, pageSize: number) => {
  return http.get<ListResponse>('/api/v1/parse/lists', {
    page_no: pageNo,
    page_size: pageSize,
  })
}
