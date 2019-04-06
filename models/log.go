package models

// 系统日志模型
type Log struct {
	TaskID    string `json:"task_id"`
	WorkerID  string `json:"worker_id"`
	UserID    string `json:"user_id"`
	Leve      string `json:"leve"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
