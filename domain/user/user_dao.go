package user

import (
	"context"
	"fmt"

	"github.com/go_bootcamp/clients/mongodb"
	"github.com/go_bootcamp/utils/date_utils"
	"github.com/go_bootcamp/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (user *User) Get() *errors.RestErr {
	filter := bson.D{primitive.E{Key: "_id", Value: user.Id}}
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	collection := client.Database(mongodb.DB).Collection(mongodb.USERS)
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}
	return nil
}

func (user *User) GetUserByEmail() *errors.RestErr {
	filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	collection := client.Database(mongodb.DB).Collection(mongodb.USERS)
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("El usuario %v no existe", user.Email))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	collection := client.Database(mongodb.DB).Collection(mongodb.USERS)

	user.CreatedAt = date_utils.GetNowString()
	user.UpdatedAt = date_utils.GetNowString()

	_, err = collection.InsertOne(context.TODO(), user)

	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	user.UpdatedAt = date_utils.GetNowString()

	update := bson.M{
		"$set": user,
	}

	collection := client.Database(mongodb.DB).Collection(mongodb.USERS)

	_, err = collection.UpdateByID(context.TODO(), user.Id, update)

	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	delete := bson.D{primitive.E{Key: "_id", Value: user.Id}}

	collection := client.Database(mongodb.DB).Collection(mongodb.USERS)

	_, err = collection.DeleteOne(context.TODO(), delete)

	if err != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Error %v", err))
	}

	return nil
}
