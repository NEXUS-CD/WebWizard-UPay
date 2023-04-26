package services

import (
	"github.com/NEXUS-CD/UPay/models"
)

func FindUserByID(id string) (*models.User, error) {
	return models.FindUserByID(id)
}
