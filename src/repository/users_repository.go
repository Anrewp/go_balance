package repository

import (
	"errors"

	"github.com/Anrewp/go_balance/src/helper"
	"github.com/Anrewp/go_balance/src/model"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(users model.Users)
	Update(users model.Users)
	Delete(userId int)
	FindById(userId int) (users model.Users, err error)
	FindAll() []model.Users
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (t *UsersRepositoryImpl) Create(users model.Users) {
	result := t.Db.Create(&users)
	helper.ErrorPanic(result.Error)

}

func (t *UsersRepositoryImpl) Update(users model.Users) {
	// var updateUser = request.UpdateUsersRequest{
	// 	Id:    users.Id,
	// 	Email: users.Email,
	// }
	result := t.Db.Model(&users).Updates(users)
	helper.ErrorPanic(result.Error)
}

func (t *UsersRepositoryImpl) Delete(usersId int) {
	var users model.Users
	result := t.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

func (t *UsersRepositoryImpl) FindById(usersId int) (model.Users, error) {
	var user model.Users
	result := t.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (t *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := t.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}
