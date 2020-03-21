package user

import (
	"net/http"
	"strconv"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-go/oauth"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/users-api/domain/users"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/users-api/service"
)

func marshallUser(user *users.User, c *gin.Context) interface{} {
	return user.Marshal(oauth.IsPublic(c.Request))
}

func marshallUsers(users users.Users, c *gin.Context) []interface{} {
	return users.Marshal(oauth.IsPublic(c.Request))
}

//getUserID Get User ID
func getUserID(userIDParam string) (int64, errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		err := errors.BadRequestError("user id sould be a number", userErr)
		return 0, err
	}
	return userID, nil
}

//Create a new User
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, saveErr := service.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, marshallUser(result, c))
}

//Get User by ID (user_id)
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	//Check the caller ID, for autorization
	// if callerID := oauth.GetCallerID(c.Request); callerID == 0 {
	// 	err := errors.RestErr{
	// 		Status:  http.StatusUnauthorized,
	// 		Message: "resource not available",
	// 		Error:   "not_authorized",
	// 	}
	// 	c.JSON(err.Status, err)
	// 	return
	// }

	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	result, getErr := service.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	if oauth.GetCallerID(c.Request) == result.ID {
		c.JSON(http.StatusOK, result.Marshal(false))
		return
	}

	c.JSON(http.StatusOK, marshallUser(result, c))
}

//Update User
func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status(), restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := service.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, marshallUser(result, c))
}

//Delete a User
func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := service.UserService.DeleteUser(userID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

//Search Users
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := service.UserService.Search(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, marshallUsers(users, c))
}

//Login Users
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, err := service.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, marshallUser(user, c))
}
