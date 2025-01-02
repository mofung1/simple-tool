interface IUserInfo {
  id?: number
  nickname?: string
  avatar?: string
  openid?: string
  unionid?: string
  mobile?: string
  gender?: number
  country?: string
  province?: string
  city?: string
  login_time?: string
  login_ip?: string
}

interface ILoginResult {
  code: number
  data: {
    token: string
    userInfo: IUserInfo
  }
  message: string
}
