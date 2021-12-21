package users

import (
	"strconv"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/domain/users"
	"github.com/delaram-gholampoor-sagha/bookstore-users-api/utils/errors"

	"net/http"

	"github.com/delaram-gholampoor-sagha/bookstore-users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// solution 1
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO : handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO handle json error
	// 	return
	// }

	// solution 2
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO handle json error
		restErr := errors.NewBadRequessrError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO  handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequessrError("user id should be a number")
		c.JSON(err.Status, err)
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
