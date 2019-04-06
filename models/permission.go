package models

// 权限模型
type (
	Permission struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
