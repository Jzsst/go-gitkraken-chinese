package utils

import "os"

func isExist(path string) {
	file, er := os.Open(path)
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()
	if er != nil && os.IsNotExist(er) {
		file, _ = os.Create(path)
		//然后将初始模板写入到zh.json
	}
}
