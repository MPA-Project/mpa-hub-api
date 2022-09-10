package universal

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PaginationQuery(c *fiber.Ctx, allowedOrderBy []string) (int, int, string, string, string, error) {
	qPageSize := c.Query("pageSize", "10")
	qPageIndex := c.Query("pageIndex", "0")
	qOrderBy := c.Query("orderBy", "createdAt")
	qOrderByDirection := c.Query("orderByDirection", "desc")
	qQuery := c.Query("query", "")

	qPageSizeInt, err := strconv.Atoi(qPageSize)
	if err != nil {
		return 0, 0, "", "", "", err
	}
	if !intInSlice(qPageSizeInt, []int{10, 25}) {
		qPageSizeInt = 10
	}

	qPageIndexInt, err := strconv.Atoi(qPageIndex)
	if err != nil {
		return 0, 0, "", "", "", err
	}
	if qPageIndexInt < 0 {
		qPageIndexInt = 0
	}
	qPageIndexInt = qPageIndexInt * qPageSizeInt

	if !stringInSlice(qOrderByDirection, []string{"asc", "desc"}) {
		qOrderByDirection = "asc"
	}

	if len(allowedOrderBy) == 0 || !stringInSlice(qOrderBy, []string{"created_at"}) {
		allowedOrderBy = append(allowedOrderBy, "created_at")
	}
	if !stringInSlice(qOrderBy, allowedOrderBy) {
		qOrderBy = "created_at"
	}

	return qPageSizeInt, qPageIndexInt, qOrderBy, qOrderByDirection, qQuery, err
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
