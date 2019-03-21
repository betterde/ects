package models

type (
	Role struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at"`
	}
)