package common

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"time"

	"myponyasia.com/hub-api/exception"
)

func RemoveExpiredFiles() {
	files, err := ReadDir("./uploads")
	exception.PanicLogging(err)

	for _, f := range files {
		if f.Name() != ".gitignore" && !f.IsDir() {
			if f.ModTime().Add(2 * time.Hour).Before(time.Now()) {
				if err := os.Remove(fmt.Sprintf("./uploads/%s", f.Name())); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func ReadDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
