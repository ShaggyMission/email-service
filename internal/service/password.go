package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"email-service/internal/config"
	"email-service/internal/models"
)

func GenerateRandomPassword() (string, error) {
	length := 9
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func UpdateUserPasswordByEmail(email string) (string, error) {
	var user models.User
	result := config.DB.Unscoped().Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found")
	}

	newPass, err := GenerateRandomPassword()
	if err != nil {
		return "", err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPass), 10)
	if err != nil {
		return "", err
	}

	user.Password = string(hashed)
	if err := config.DB.Save(&user).Error; err != nil {
		return "", err
	}

	return newPass, nil
}

