package enums

func StatusEnum(status string) int {
	switch status {
	case "ACTIVE":
		return 1
	case "INACTIVE":
		return 2
	case "DELETED":
		return 0
	default:
		return 1
	}
}
