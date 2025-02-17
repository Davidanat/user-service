package seeder

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	users := []models.User{
		{
			UUID:        uuid.New(),
			Username:    "admin",
			Name:        "Administrator",
			Password:    string(password),
			PhoneNumber: "08123456789",
			Email:       "admin@t4t.org",
			RoleID:      constants.Admin,
		},
		{
			UUID:        uuid.New(),
			Username:    "customer",
			Name:        "Customer",
			Password:    string(password),
			PhoneNumber: "08123456789",
			Email:       "user@t4t.org",
			RoleID:      constants.Customer,
		},
	}

	for _, user := range users {
		err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
		if err != nil {
			logrus.Errorf("failed to seed user: %v", err)
			panic(err)
		}
		logrus.Infof("user %s seeded", user.Username)
	}
}
