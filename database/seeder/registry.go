package seeder

import "gorm.io/gorm"

type Resistry struct {
	db *gorm.DB
}

type ISeederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) ISeederRegistry {
	return &Resistry{db: db}
}

func (r *Resistry) Run() {
	RunRoleSeeder(r.db)
	RunUserSeeder(r.db)
}
