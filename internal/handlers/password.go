package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"email-service/internal/service"
	"email-service/pkg/mail"
)

type RecoveryRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func RecoverPasswordHandler(c *gin.Context) {
	var req RecoveryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
		return
	}

	newPassword, err := service.UpdateUserPasswordByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found or could not update password"})
		return
	}

	if err := mail.SendRecoveryEmail(req.Email, newPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send recovery email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully. Please check your email."})
}
