package message

func userMessage() map[string]map[string]string {
	return map[string]map[string]string{
		"Name": {
			"required": "请填写用户姓名",
		},
		"Email": {
			"email": "请填写正确的邮箱地址",
		},
		"Pass": {
			"required": "请填写用户密码",
		},
		"Confirm": {
			"eqfield": "请确认用户密码",
		},
		"TeamID": {
			"uuid4": "团队ID有误",
		},
		"Manager": {
			"required": "请选择是否是管理员",
		},
	}
}
