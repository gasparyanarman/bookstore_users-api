package users

import (
	"net/http"
	"strconv"

	"github.com/gasparyanarman/bookstore_users-api/domain/users"
	"github.com/gasparyanarman/bookstore_users-api/services"
	"github.com/gasparyanarman/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)

	if userErr != nil {
		return 0, errors.NewBadRequestError("User id should be a number")
	}

	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	id, err := getUserId(c.Param("user_id"))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(id)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	updatedUser, updateErr := services.UpdateUser(isPartial, &user)
	if updateErr != nil {
		upErr := errors.NewBadRequestError("Update failed")
		c.JSON(updateErr.Status, upErr)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(userErr.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
