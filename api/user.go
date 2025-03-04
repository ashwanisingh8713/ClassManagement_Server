package api

import (
	"ClassManagement/database"
	"ClassManagement/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// changePassword godoc
// @Summary Change Password of the user
// @Description To Change Password by existing user
// @Tags User Module
// @Produce json
// @Success 200 {object} StatusInfoResponse
// @Param       json  body ChangePasswordInput true "It takes json as input"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /changePassword [post]
func changePassword(c *gin.Context) {
	var changePasswordInput ChangePasswordInput
	if err := c.ShouldBindJSON(&changePasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	var statusInfo StatusInfoResponse
	_, err := database.GetUserByEmailAndPasswordHash(changePasswordInput.Email, changePasswordInput.OldPassword)
	if err != nil {
		statusInfo = StatusInfoResponse{Msg: "Email or Password is incorrect", Code: http.StatusAccepted, Status: false}
		c.JSON(http.StatusAccepted, gin.H{"data": statusInfo})
	}
	err = database.UpdatePasswordByEmail(changePasswordInput.Email, changePasswordInput.NewPassword)
	if err != nil {
		statusInfo = StatusInfoResponse{Msg: "Password is not updated", Code: http.StatusAccepted, Status: false}
		c.JSON(http.StatusAccepted, gin.H{"data": statusInfo})
	} else {
		statusInfo = StatusInfoResponse{Msg: "Password is updated", Code: http.StatusAccepted, Status: true}
		c.JSON(http.StatusOK, gin.H{"data": statusInfo})
	}
}

// forgotPassword godoc
// @Summary Forgot Password
// @Description To retrieve the password
// @Tags User Module
// @Accept json
// @Produce json
// @Success 200 {object} database.User
// @Param       json  body ForgotPasswordInput true "It takes json as input"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /forgotPassword [post]
func forgotPassword(c *gin.Context) {

}

// isUserExist godoc
// @Summary To check whether Email already registered with any User or not.
// @Description To check if user exists or not
// @Tags User Module
// @Produce json
// @Success 200 {object} EmailExistResponse
// @Param       json  body EmailInput true "It takes email as input"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /isUserExist [post]
func isUserExist(c *gin.Context) {
	var emailInput EmailInput
	if err := c.ShouldBindJSON(&emailInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var emailExist EmailExistResponse
	isExist, err := database.IsUserExist(emailInput.Email)
	if err != nil {
		emailExist = EmailExistResponse{Email: emailInput.Email, Msg: err.Error(), IsExist: isExist, Code: http.StatusBadRequest}
	} else {
		emailExist = EmailExistResponse{Email: emailInput.Email, Msg: "Email is available to create new account.", IsExist: isExist, Code: http.StatusOK}
	}
	c.JSON(http.StatusOK, gin.H{"data": emailExist})
}

// signIn godoc
// @Summary Sign In user
// @Description To login by existing user
// @Tags User Module
// @Produce json
// @Success 200 {object} database.User
// @Param       json  body EmailPassword true "It takes json as input"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /signIn [post]
func signIn(c *gin.Context) {
	var userReq EmailPassword
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := database.GetUserByEmailAndPasswordHash(userReq.Email, userReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Password is incorrect"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"data": user})
	}
}

// signUp godoc
// @Summary Sign Up by User
// @Description To create a new user
// @Tags User Module
// @Produce json
// @Success 200 {object} model.CreateUser
// @Param       json  body model.CreateUser true "It takes json as input"
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

type ChangePasswordInput struct {
	Email       string `json:"email,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type ForgotPasswordInput struct {
	Email string `json:"email,omitempty"`
}

type EmailExistResponse struct {
	Email   string `json:"email,omitempty"`
	Msg     string `json:"msg,omitempty"`
	IsExist bool   `json:"isExist,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type EmailInput struct {
	Email string `json:"email,omitempty"`
}

type EmailPassword struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type StatusInfoResponse struct {
	Msg    string `json:"msg,omitempty"`
	Code   int    `json:"code,omitempty"`
	Status bool   `json:"status,omitempty"`
}
