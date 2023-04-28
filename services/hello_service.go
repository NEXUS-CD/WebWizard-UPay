package services

import (
	"UPay/models"
)

func FindUserByID(id string) (*models.User, error) {
	return models.FindUserByID(id)
}
