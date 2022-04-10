package users

import (
	"os/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserById(id primitive.ObjectID) (*user.User, *errors.RestErr) {
	result := &user.User{Id: id}
	if err := result.Get(); err != nil {
		return result, err
	}
	return result, nil
}

func GetUserByEmail(email string) (*user.User, *errors.RestErr) {
	result := &user.User{Email: email}
	if err := result.GetUserByEmail(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user user.User) (*user.User, *errors.RestErr) {
	err := user.GetUserByEmail()
	if err == nil {
		return nil, errors.NewBadRequestError("E-mail" + user.Email + " is already been registered")
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user user.User) (*user.User, *errors.RestErr) {
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(user user.User) (bool, *errors.RestErr) {
	if err := user.Delete(); err != nil {
		return false, err
	}

	return true, nil
}
