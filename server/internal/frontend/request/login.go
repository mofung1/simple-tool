package request

type MnpLoginRequest struct {
	Code     string `json:"code" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar" `
	Gender   int    `json:"gender"`
}
