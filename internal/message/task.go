package message

func taskMessage() map[string]map[string]string {
	return map[string]map[string]string{
		"Name": {
			"required": "请填写任务名称",
		},
		"ParentID": {
			"uuid4": "父任务ID有误",
		},
		"Content": {
			"required": "请填写任务内容",
		}, "Event": {
			"required": "请选择触发事件",
		}, "Mode": {
			"required": "请选择任务类型",
		}, "Overlap": {
			"required": "请选择是否重复执行",
		}, "Timeout": {
			"gte": "请填写超时时间",
		}, "Interval": {
			"gte": "请填写重试间隔时间",
		}, "Retries": {
			"gte": "请填写重试次数",
		}, "Status": {
			"required": "请选择任务状态",
		},
	}
}
