package interfaces

import (
	"errors"

	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"

	"gorm.io/gorm"
)

type UsersInterfaceImpl struct {
	Db *gorm.DB
}

func NewUsersInterfaceImpl(Db *gorm.DB) UsersInterface {
	return &UsersInterfaceImpl{Db: Db}
}

// Delete implements UsersInterface
func (user *UsersInterfaceImpl) Delete(userId int) {
	var users model.User
	result := user.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersInterface
func (user *UsersInterfaceImpl) FindAll() []model.User {
	var users []model.User
	result := user.Db.Preload("Posts").Find(&users)
	// result := user.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindByEmail implements UsersInterface
// func (user *UsersInterfaceImpl) FindByEmail(email string) (model.User, error) {
// 	var users model.User
// 	result := user.Db.Preload("Posts").First(&users, userId)
// 	if result.Error != nil {
// 		return users, errors.New("user not found")
// 	}
// 	return users, nil
// }

// FindById implements UsersInterface
func (user *UsersInterfaceImpl) FindById(userId int) (model.User, error) {
	var users model.User
	result := user.Db.First(&users, userId)
	if result.Error != nil {
		return users, errors.New("user not found")
	}
	return users, nil
}

// FindByUsername implements UsersInterface
func (user *UsersInterfaceImpl) FindByUsername(username string) (model.User, error) {
	var users model.User
	result := user.Db.First(&users, "username = ?", username)
	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// FindByEmail implements UsersInterface.
func (user *UsersInterfaceImpl) FindByEmail(email string) (model.User, error) {
	var users model.User
	result := user.Db.First(&users, "email = ?", email)
	if result.Error != nil {
		return users, errors.New("invalid email or password")
	}
	return users, nil
}

// Save implements UsersInterface
func (user *UsersInterfaceImpl) Save(users model.User) {
	result := user.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersInterface
func (user *UsersInterfaceImpl) Update(users model.User) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}

	result := user.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}
