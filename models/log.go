package models

// 系统日志模型
type Log struct {
	ID        int    `json:"id" xorm:"pk autoincr notnull 'id'"`
	TaskID    string `json:"task_id" xorm:"char(36) default '' 'task_id'"`
	WorkerID  string `json:"worker_id" xorm:"char(36) default '' 'worker_id'"`
	UserID    string `json:"user_id" xorm:"char(36) default '' 'user_id'"`
	Leve      string `json:"leve" xorm:"varchar(32) notnull"`
	Status    string `json:"status" xorm:"varchar(32) notnull default 'running'"`
	Content   string `json:"content" xorm:""`
	CreatedAt string `json:"created_at" xorm:"datetime notnull created"`
}
