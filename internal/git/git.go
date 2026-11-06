package git

import (
	"os"
)

func GetGitDir(dir string) (bool, error) {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, v := range dirContent {
		if v.IsDir() && v.Name() == ".git" {
			return true, nil
		}
	}
	return false, nil
}
