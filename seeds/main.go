package seeds

import (
	"github.com/betterde/ects/models"
	"log"
)

type (
	SeederInterface interface {
		Register(seeder *Seeder)
	}

	Seeder struct {
		instances []models.Seeder
	}
)

func (seeder *Seeder) Run() error {
	seeders := []SeederInterface{
		&MenuSeeder{},
		&PermissionSeeder{},
	}

	for _, item := range seeders {
		item.Register(seeder)
	}

	for _, instalce := range seeder.instances {
		if err := instalce.Seed(); err != nil {
			log.Println(err)
		}
	}
	return nil
}
