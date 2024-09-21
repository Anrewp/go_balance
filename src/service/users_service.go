package service

import (
	"github.com/Anrewp/go_balance/src/data/request"
	"github.com/Anrewp/go_balance/src/data/response"
	"github.com/Anrewp/go_balance/src/helper"
	"github.com/Anrewp/go_balance/src/model"
	"github.com/Anrewp/go_balance/src/repository"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(usersId int) response.UsersResponse
	FindAll() []response.UsersResponse
}

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewUsersServiceImpl(usersRepository repository.UsersRepository) UsersService {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
	}
}

func (t *UsersServiceImpl) Create(user request.CreateUsersRequest) {
	// err := t.Validate.Struct(user)
	// helper.ErrorPanic(err)
	userModel := model.Users{
		Email: user.Email,
	}
	t.UsersRepository.Create(userModel)
}

func (t *UsersServiceImpl) Update(user request.UpdateUsersRequest) {
	usersData, err := t.UsersRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	usersData.Email = user.Email
	t.UsersRepository.Update(usersData)
}

func (t *UsersServiceImpl) Delete(userId int) {
	t.UsersRepository.Delete(userId)
}

func (t *UsersServiceImpl) FindById(userId int) response.UsersResponse {
	userData, err := t.UsersRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:    userData.Id,
		Email: userData.Email,
	}

	return userResponse
}

func (t *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:    value.Id,
			Email: value.Email,
		}
		users = append(users, user)
	}
	return users
}
