package model

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Username string `gorm:"type:varchar(100);not null" json:"username" form:"username"`
}

func CreateUser(db *gorm.DB, user User) error {
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadUser(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, id int, user User) (*User, error) {
	var existingUser User
	if result := db.First(&existingUser, id); result.Error != nil {
		return nil, result.Error
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}

	if result := db.Save(&existingUser); result.Error != nil {
		return nil, result.Error
	}
	return &existingUser, nil
}

func DeleteUser(db *gorm.DB, id int) error {
	var user User
	if result := db.First(&user, id); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
