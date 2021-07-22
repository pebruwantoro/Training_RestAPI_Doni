package database

import (
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/middlewares"
	"pebruwantoro/middleware/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//jwt middlewares
func LoginUsers(email, password string) (interface{}, error) {
	var user models.User
	var err error
	if err = config.DB.Where("email = ? AND password =?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

func GetDetailUsers(userId int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteOneUser(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateOneUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
