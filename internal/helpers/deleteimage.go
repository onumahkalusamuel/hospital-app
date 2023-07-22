package helpers

import (
	"os"
	"path"
)

func DeleteImage(folder string, source string) (string, bool) {
	filename := path.Join(folder, source)
	if err := os.Remove(filename); err != nil {
		return err.Error(), false
	}
	return "", true
}
