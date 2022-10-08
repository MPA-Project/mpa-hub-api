package universal

import (
	"fmt"
	"os"

	"myponyasia.com/hub-api/pkg/entities"
)

func GenerateCDNUrl(basePath string, fileManagerModelImage *entities.FileManagerModelImage) string {
	STORAGE_URL := os.Getenv("STORAGE_PUBLIC_URL")
	return fmt.Sprintf("%s%s/%s/%s/%s/%s", STORAGE_URL, basePath, fileManagerModelImage.PYear, fileManagerModelImage.PMonth, fileManagerModelImage.PDay, fileManagerModelImage.Filename)
}
