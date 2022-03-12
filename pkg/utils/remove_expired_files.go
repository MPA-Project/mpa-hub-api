package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func RemoveExpiredFiles() {
	files, err := ioutil.ReadDir("./uploads")
	if err != nil {
		fmt.Println(err)
	}

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
