package response

import "simple-tool/server/internal/models"

type LoginResult struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
