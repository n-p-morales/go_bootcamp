package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_bootcamp/services/users"
	"github.com/n-p-morales/go_bootcamp/domain/user"
	"github.com/n-p-morales/go_bootcamp/services/users"
	"github.com/n-p-morales/go_bootcamp/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Get user by email function
func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	res, getErr := users.GetUserByEmail(email)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

//Create user function
func CreateUser(c *gin.Context) {
	var user user.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := users.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//Update user function
func UpdateUser(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Param("_id"))

	if err != nil {
		restErr := errors.NewBadRequestError("Error convert string to Object")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, getErr := users.GetUserById(_id)
	if getErr != nil {
		restErr := errors.NewBadRequestError("User does exist")
		c.JSON(restErr.Status, restErr)
		return
	}
	var user user.User

	user.Id = res.Id

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, updateErr := users.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//Delete user function
func DeleteUser(c *gin.Context) {
	_id, err := primitive.ObjectIDFromHex(c.Param("_id"))

	if err != nil {
		restErr := errors.NewBadRequestError("Error convert string to object")
		c.JSON(restErr.Status, restErr)
	}

	res, getErr := users.GetUserById(_id)
	if getErr != nil {
		restErr := errors.NewBadRequestError("User doesn't exist")
		c.JSON(restErr.Status, restErr)
		return
	}
	var user user.User

	user.Id = res.Id

	result, deleteUser := users.DeleteUser(user)

	if deleteUser != nil {
		c.JSON(deleteUser.Status, deleteUser)
		return
	}
	c.JSON(http.StatusOK, result)
}
