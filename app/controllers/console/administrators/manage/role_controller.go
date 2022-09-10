package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/app/services/universal"
	"myponyasia.com/hub-api/pkg/database"
)

type RolesResponse struct {
	ID    uuid.UUID
	Name  string
	Level int

	UserCount int
}

func RoleList(c *fiber.Ctx) error {

	qPageSize, qPageIndex, qOrderBy, qOrderByDirection, qQuery, err := universal.PaginationQuery(c, []string{"name"})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var dataCount int64
	dataCountQuery := database.DB.Model(&models.Role{})
	if len(qQuery) >= 3 {
		dataCountQuery = dataCountQuery.Where("name LIKE ?", "%"+qQuery+"%")
	}
	if err := dataCountQuery.Count(&dataCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var data []models.Role
	dataQuery := database.DB.Model(&models.Role{})
	if len(qQuery) >= 3 {
		dataQuery = dataQuery.Where("name LIKE ?", "%"+qQuery+"%")
	}
	if dataCount > 0 {
		if err := dataQuery.Order(qOrderBy + " " + qOrderByDirection).Limit(qPageSize).Offset(qPageIndex).Find(&data).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var dataResponse []RolesResponse = []RolesResponse{}
	for _, role := range data {
		dataResponse = append(dataResponse, RolesResponse{
			ID:        role.ID,
			Name:      role.Name,
			Level:     role.Level,
			UserCount: 0,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"list":  dataResponse,
			"total": dataCount,
		},
		"dbg": fiber.Map{
			"pageSize":         qPageSize,
			"pageIndex":        qPageIndex,
			"orderBy":          qOrderBy,
			"orderByDirection": qOrderByDirection,
			"query":            qQuery,
			"queryLength":      len(qQuery),
		},
	})
}
