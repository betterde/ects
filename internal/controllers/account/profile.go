package account

type (
	ProfileController struct {
	}
	Profile struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Avatar  string `json:"avatar"`
		GroupId int64  `json:"group_id"`
		Manager bool   `json:"manager"`
	}
)
