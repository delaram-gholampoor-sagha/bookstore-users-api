package users

import (
	"strconv"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"

	"net/http"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/services"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequessrError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	// solution 1
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO : handle error
	// 	return
	// }
	// the marshal takes the json input and tries to populate the given struct
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO handle json error
	// 	return
	// }

	// solution 2
	// take the parameters that you need to process that request
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO handle json error
		restErr := errors.NewBadRequessrError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// the controller is not in charge of knowing  where and how we are storing different users
	//and send that request to the service
	result, saveErr := services.CreateUser(user)
	// if have any error we are gonna return that error and return
	if saveErr != nil {
		//TODO  handle user creation error
		// this is the status code and this is the json response
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {

		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO  handle user creation error
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	// take the user id
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {

		c.JSON(idErr.Status, idErr)
		return
	}

	// take the json body
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO handle json error
		restErr := errors.NewBadRequessrError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	// take the user id
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {

		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}
