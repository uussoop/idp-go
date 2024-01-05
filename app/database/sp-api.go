package database

import "gorm.io/gorm"

func (s *ServiceProviders) Create() (err error) {
	err = DB.Create(&s).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceProviders) Update() (err error) {
	return DB.Model(s).Updates(s).Error
}

func GetAllServiceProviders() (s []ServiceProviders, err error) {
	err = DB.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}
func GetByIpAndToken(ip, token string) (err error) {
	var sp *ServiceProviders
	err = DB.Where(ServiceProviders{Ip: token, Token: token}).First(&sp).Error
	if err != nil {
		return err
	}
	if sp.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	return
}
