package services

import "github.com/JuanFernandez87/go-vue-login-with-google/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
	UpsertUser(string, *models.UpdateDBUser) (*models.DBResponse, error)
}
