package service

import (
	"gitlab.com/aubayaml/aubayaml-go/bookstore/users-api/domain/users"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/users-api/util/date_utils"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/crypto"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	//UserService encapsulado
	UserService userServicesInterface = &userServices{}
)

type userServices struct{}

type userServicesInterface interface {
	GetUser(int64) (*users.User, errors.RestErr)
	CreateUser(users.User) (*users.User, errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, errors.RestErr)
	DeleteUser(int64) errors.RestErr
	Search(string) (users.Users, errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, errors.RestErr)
}

//GetUser Get user by ID
func (s *userServices) GetUser(userID int64) (*users.User, errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser Service to persist User
func (s *userServices) CreateUser(user users.User) (*users.User, errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowAsDataBaseString()
	user.Password = crypto.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser current user
func (s *userServices) UpdateUser(isPartial bool, user users.User) (*users.User, errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	user.Trim()
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

//DeleteUser remove a given user
func (s *userServices) DeleteUser(userID int64) errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

//Search return users given status
func (s *userServices) Search(status string) (users.Users, errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

//Search return users given status
func (s *userServices) LoginUser(request users.LoginRequest) (*users.User, errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
		Status:   users.StatusActive,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
