package seeds

import "github.com/betterde/ects/models"

type PermissionSeeder struct {
}

// 注册数据
func (permission *PermissionSeeder) Register(seeder *Seeder) {
	items := []models.Seeder{
		&models.Permission{
			Code:        "dashboard",
			Name:        "概览",
			Module:      "dashboard",
			Description: "概览页面权限",
		},
		&models.Permission{
			Code:        "task_index",
			Name:        "查看任务",
			Module:      "task",
			Description: "查看任务列表的权限",
		},
		&models.Permission{
			Code:        "task_create",
			Name:        "创建任务",
			Module:      "task",
			Description: "创建任务的权限",
		},
		&models.Permission{
			Code:        "task_update",
			Name:        "更新任务",
			Module:      "task",
			Description: "修改任务属性的权限",
		},
		&models.Permission{
			Code:        "task_delete",
			Name:        "删除任务",
			Module:      "task",
			Description: "删除任务的权限",
		},
		&models.Permission{
			Code:        "worker_index",
			Name:        "查看节点",
			Module:      "worker",
			Description: "查看节点列表的权限",
		},
		&models.Permission{
			Code:        "worker_create",
			Name:        "创建节点",
			Module:      "worker",
			Description: "创建节点的权限",
		},
		&models.Permission{
			Code:        "worker_update",
			Name:        "更新节点",
			Module:      "worker",
			Description: "更新节点属性的权限",
		},
		&models.Permission{
			Code:        "worker_delete",
			Name:        "删除节点",
			Module:      "worker",
			Description: "删除节点的权限",
		},
		&models.Permission{
			Code:        "team_index",
			Name:        "查看团队",
			Module:      "team",
			Description: "查看团队列表的权限",
		},
		&models.Permission{
			Code:        "team_create",
			Name:        "创建团队",
			Module:      "team",
			Description: "创建团队的权限",
		}, &models.Permission{
			Code:        "team_update",
			Name:        "更新团队",
			Module:      "team",
			Description: "更新团队属性的权限",
		},
		&models.Permission{
			Code:        "team_delete",
			Name:        "删除团队",
			Module:      "team",
			Description: "删除团队的权限",
		},
		&models.Permission{
			Code:        "user_index",
			Name:        "查看用户",
			Module:      "user",
			Description: "查看用户列表的权限",
		},
		&models.Permission{
			Code:        "user_create",
			Name:        "创建用户",
			Module:      "user",
			Description: "创建用户的权限",
		},
		&models.Permission{
			Code:        "user_update",
			Name:        "更新用户",
			Module:      "user",
			Description: "更新用户属性的权限",
		},
		&models.Permission{
			Code:        "user_delete",
			Name:        "删除用户",
			Module:      "user",
			Description: "删除用户的权限",
		},
		&models.Permission{
			Code:        "pipeline_index",
			Name:        "查看流水线",
			Module:      "pipeline",
			Description: "查看流水线列表的权限",
		},
		&models.Permission{
			Code:        "pipeline_create",
			Name:        "创建流水线",
			Module:      "pipeline",
			Description: "创建流水线的权限",
		},
		&models.Permission{
			Code:        "pipeline_update",
			Name:        "更新流水线",
			Module:      "pipeline",
			Description: "更新流水线属性的权限",
		},
		&models.Permission{
			Code:        "pipeline_delete",
			Name:        "删除流水线",
			Module:      "pipeline",
			Description: "删除流水线的权限",
		},
		&models.Permission{
			Code:        "record_index",
			Name:        "查看记录",
			Module:      "record",
			Description: "查看任务执行记录的权限",
		},
		&models.Permission{
			Code:        "record_clear",
			Name:        "清理记录",
			Module:      "record",
			Description: "清理任务执行记录的权限",
		},
		&models.Permission{
			Code:        "log_index",
			Name:        "查看日志",
			Module:      "log",
			Description: "查看用户操作记录的权限",
		},
		&models.Permission{
			Code:        "setting_index",
			Name:        "查看设置",
			Module:      "setting",
			Description: "查看系统设置的权限",
		},
		&models.Permission{
			Code:        "setting_update",
			Name:        "修改设置",
			Module:      "setting",
			Description: "更新系统设置的权限",
		},
		&models.Permission{
			Code:        "setting_upgrade",
			Name:        "升级系统",
			Module:      "setting",
			Description: "升级系统版本的权限",
		},
	}

	seeder.instances = append(seeder.instances, items...)
}
