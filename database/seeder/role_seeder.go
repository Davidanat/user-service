package seeder

import (
	"user-service/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunRoleSeeder(db *gorm.DB) {
	roles := []models.Role{
		{
			Code: "ADMIN",
			Name: "Administrator",
		},
		{
			Code: "CUSTOMER",
			Name: "Customer",
		},
	}

	for _, role := range roles {
		err := db.FirstOrCreate(&role, models.Role{Code: role.Code}).Error
		if err != nil {
			logrus.Errorf("failed to seed role: %v", err)
			panic(err)
		}
		logrus.Infof("role %s seeded", role.Code)
	}
}
