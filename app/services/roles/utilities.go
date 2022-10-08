package roles

import (
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/entities"
)

func FindRolesByName(name string) (entities.Role, error) {
	var roles entities.Role
	err := database.DB.Where("name = ?", name).First(&roles).Error
	if err != nil {
		return entities.Role{}, err
	}
	return roles, nil
}

func FindRolesByUserId(user entities.User) ([]entities.UserRoles, error) {
	var userRoles []entities.UserRoles
	err := database.DB.Where("user_id = ?", user.ID).Preload("Role").Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
