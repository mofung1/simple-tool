package request

type MnpLoginRequest struct {
	Code string `json:"code" binding:"required"`
}
