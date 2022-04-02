package roles

import (
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
)

func FindRolesByName(name string) (models.Role, error) {
	var roles models.Role
	err := database.DB.Where("name = ?", name).First(&roles).Error
	if err != nil {
		return models.Role{}, err
	}
	return roles, nil
}

func FindRolesByUserId(user models.User) ([]models.UserRoles, error) {
	var userRoles []models.UserRoles
	err := database.DB.Where("user_id = ?", user.ID).Preload("Role").Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
