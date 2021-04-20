package database

import (
	"github.com/blyndusk/cofy/api/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	_ = Db.AutoMigrate(&models.User{})
	logrus.Info("Migrations done !")
	LoadFakeData()

}
func LoadFakeData() {

	var users = []models.User{
		{Name: "cofy1", Coins: 100, Xp: 123},
	}
	Db.Create(&users)
	logrus.Info("Fake data loaded !")

}
