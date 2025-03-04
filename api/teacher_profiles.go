package api

import (
	"ClassManagement/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createTeacherProfile godoc
// @Summary Create Teacher Profile
// @Description To Create Teacher Profile
// @Tags Teacher
// @Accept json
// @Produce json
// @Success 200 {object} database.TeacherProfile
// @Param       json  body database.TeacherProfile true "It takes json as input"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createTeacherProfile [post]
func createTeacherProfile(c *gin.Context) {
	var teacherProfile database.TeacherProfile
	if err := c.ShouldBindJSON(&teacherProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if teacherProfile.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "UserID is required"})
		return
	}
	err := database.CreateTeacherProfile(teacherProfile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": teacherProfile})
	}
}

// getTeacherProfile godoc
// @Summary Retrieve Teacher Profile
// @Description To Retrieve Teacher Profile
// @Tags Teacher
// @Accept json
// @Produce json
// @Success 200 {object} database.TeacherProfile
// @Param user_id path int true "User ID"
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getTeacherProfile/{user_id} [post]
func getTeacherProfile(c *gin.Context) {
	userID := c.Param("user_id")
	profile, err := database.GetTeacherProfile(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": profile})
	}
}
