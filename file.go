package file

import (
	"os"
	"time"
)

func GetFileCreatedTime(filename string) time.Time {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}
	}
	return fileInfo.ModTime()
}
