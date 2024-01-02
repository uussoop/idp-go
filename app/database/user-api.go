package database

import "gorm.io/gorm"

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

	if user.ID == 0 || !user.Verified {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil

}
