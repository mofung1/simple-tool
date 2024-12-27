interface IUserInfo {
  nickname?: string
  avatar?: string
  token?: string
  openid?: string
  mobile?: string
  gender?: number
  country?: string
  province?: string
  city?: string
}

interface ILoginResult {
  code: number
  data: {
    token: string
    userInfo: IUserInfo
  }
  message: string
}
