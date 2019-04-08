package seeds

import "github.com/betterde/ects/models"

type (
	MenuSeeder struct {
	}
)

func (menus *MenuSeeder) Register(seeder *Seeder) {
	items := []models.Seeder{
		&models.Menu{
			Code:           "dashboard",
			ParentCode:     "",
			PermissionCode: "dashboard",
			Name:           "概览",
			Route:          "/dashboard",
			Icon:           "",
		},
		&models.Menu{
			Code:           "task",
			ParentCode:     "",
			PermissionCode: "task",
			Name:           "任务",
			Route:          "/task",
			Icon:           "",
		},
		&models.Menu{
			Code:           "worker",
			ParentCode:     "",
			PermissionCode: "worker",
			Name:           "节点",
			Route:          "/worker",
			Icon:           "",
		},
		&models.Menu{
			Code:           "team",
			ParentCode:     "",
			PermissionCode: "team",
			Name:           "团队",
			Route:          "/team",
			Icon:           "",
		},
		&models.Menu{
			Code:           "user",
			ParentCode:     "",
			PermissionCode: "user",
			Name:           "用户",
			Route:          "/user",
			Icon:           "",
		},
		&models.Menu{
			Code:           "log",
			ParentCode:     "",
			PermissionCode: "log",
			Name:           "日志",
			Route:          "/log",
			Icon:           "",
		},
		&models.Menu{
			Code:           "record",
			ParentCode:     "",
			PermissionCode: "record",
			Name:           "记录",
			Route:          "/record",
			Icon:           "",
		},
		&models.Menu{
			Code:           "setting",
			ParentCode:     "",
			PermissionCode: "setting",
			Name:           "设置",
			Route:          "/setting",
			Icon:           "",
		},
	}

	seeder.instances = append(seeder.instances, items...)
}
