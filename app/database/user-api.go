package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (u *User) Create() (err error) {
	err = DB.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}
func (u *User) Update() (err error) {
	return DB.Model(u).Updates(u).Error
}
func GetUserByAddress(address string) (user *User, err error) {
	err = DB.Where(User{Address: address}).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil

}
func GetWhitelistByAddress(address string) bool {
	var user *UserWhitelist
	err := DB.Where(UserWhitelist{Address: address}).First(&user).Error
	logrus.Warn("get whitelist err: ", err)
	logrus.Warn("get whitelist usr: ", user)

	if err != nil {
		return false
	}

	if user.ID == 0 {
		return false
	}
	return true

}
func (u *UserWhitelist) Create() (err error) {
	err = DB.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}
