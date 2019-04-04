package models

const (
	STATUS_CONNECTED = "connected"
	STATUS_DISCONNECTED = "disconnected"
)

type (
	Worker struct {
		ID         string `json:"id" xorm:"pk char(36) notnull 'id'"`
		Name       string `json:"name" xorm:"varchar(255) notnull 'name'"`
		IP         string `json:"ip" xorm:"varchar(15) null 'ip'"`
		Status     string `json:"status" xorm:"-"`
		Remark     string `json:"remark" xorm:"varchar(255) null 'remark'"`
		CreatedAt  string `json:"created_at" xorm:"datetime notnull created"`
		UpdatedAt  string `json:"updated_at" xorm:"datetime notnull updated"`
		Model `json:"-" xorm:"-"`
	}
)

// 创建节点
func (worker *Worker) Store() error {
	_, err := Engine.Insert(worker)
	return err
}

// 更新节点信息
func (worker *Worker) Update() error {
	_, err := Engine.Id(worker.ID).Update(worker)
	return err
}