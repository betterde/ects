package models

type Menu struct {
	Code           string `xorm:"not null pk comment('编码') VARCHAR(255)"`
	ParentCode     string `xorm:"not null comment('父编码') index VARCHAR(255)"`
	PermissionCode string `xorm:"not null comment('权限编码') index VARCHAR(255)"`
	Name           string `xorm:"not null comment('名称') VARCHAR(255)"`
	Route          string `xorm:"not null comment('前端路由') VARCHAR(255)"`
	Icon           string `xorm:"not null comment('图标') VARCHAR(255)"`
}

// 定义模型的数据表名称
func (menu *Menu) TableName() string {
	return "menus"
}

// 数据填充器
func (menu *Menu) Seed() error {
	return menu.Store()
}

// 创建菜单
func (menu *Menu) Store() error {
	_, err := Engine.Insert(menu)
	return err
}
