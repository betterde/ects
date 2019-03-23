package models

import "net"

type (
	Worker struct {
		ID string `json:"id" xorm:"pk char(36) notnull 'id'"`
		Name string `json:"name" xorm:"varchar(255) notnull 'name'"`
		IP net.IP `json:"ip" xorm:"int(10) notnull 'ip'"`
		Port int `json:"port" xorm:"int notnull 'ip'"`
		Status string `json:"status" xorm:"varchar(16) notnull default 'running' 'status'"`
		Remark string `json:"remark" xorm:"varchar(255) null 'remark'"`
	}
)