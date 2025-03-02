package api

import (
	"ClassManagement/database"
	"ClassManagement/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// isUserExist godoc
// @Summary User already exist
// @Description To check if user exists or not
// @Tags User Module
// @Produce json
// @Success 200 {object} EmailCheck
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /isUserExist [post]
func isUserExist(c *gin.Context) {
	email := ""
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isExist, err := database.IsUserExist(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isExist {
		c.JSON(http.StatusOK, gin.H{"data": "User Exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "User Not Exist"})
		return
	}
}

// signUp godoc
// @Summary Signup user
// @Description To create a new user
// @Tags User Module
// @Produce json
// @Success 200 {object} model.User
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /signUp [post]
func signUp(c *gin.Context) {
	var user model.CreateUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isCreated, msg := database.CreateUser(user.Name, "passwordHash", user.Role, user.Email)
	if isCreated {
		c.JSON(http.StatusOK, gin.H{"data": msg})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	}
}

// signIn godoc
// @Summary Signup user
// @Description To create a new user
// @Tags User Module
// @Produce json
// @Success 200 {object} database.User
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /signIn [post]
func signIn(c *gin.Context) {
	var userReq model.CreateUser
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := database.GetUserByEmailAndPasswordHash(userReq.Email, userReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Password is incorrect"})
	}
}

type EmailCheck struct {
	Email string `json:"email,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
