package helpers

import (
	"os"
	"path"

	"github.com/onumahkalusamuel/hospital-app/config"
)

func DeleteImage(folder string, source string) (string, bool) {
	filename := path.Join(config.FilesFolder, path.Join(folder, source))
	if err := os.Remove(filename); err != nil {
		return err.Error(), false
	}
	return "", true
}
