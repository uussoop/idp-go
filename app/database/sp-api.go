package database

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
