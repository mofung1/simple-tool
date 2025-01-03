interface IUserInfo {
  id?: number
  sn?: number
  nickname?: string
  avatar?: string
  username?: string
  phone?: string
  gender?: number
  is_disable?: number
  login_ip?: string
  login_time?: string
}

interface ILoginResult {
  code: number
  data: {
    token: string
    user: IUserInfo
  }
  message: string
}
