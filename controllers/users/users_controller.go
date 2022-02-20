package users

import (
	"github.com/gasparyanarman/bookstore_users-api/domain/users"
	"github.com/gasparyanarman/bookstore_users-api/services"
	"github.com/gasparyanarman/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}

	c.JSON(http.StatusOK, user)
}
