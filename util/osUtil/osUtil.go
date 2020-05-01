package osUtil

import (
	"errors"
	"os"
)

func TouchFile(fileName string, path string) (*os.File, error) {
	// Check File Is Existence
	_, err := os.Stat(path+fileName)
	if err != nil {
		_, err := os.Stat(path)
		if err != nil {
			// 没有对应文件夹
			mkErr := os.Mkdir(path, 0777)
			if mkErr != nil {
				return nil, errors.New("mkdir_failed_"+path)
			}
		}

		log, tcErr := os.Create(path+fileName)
		if tcErr != nil {
			return nil, errors.New("touch_failed_"+path+fileName)
		}

		return log, nil
	} else {
		return os.OpenFile(path+fileName, os.O_APPEND|os.O_WRONLY, 0777)
	}

}
