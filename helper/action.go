package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CopyDir(src string, dest string) error {

	if dest == src {
		return fmt.Errorf("cannot copy a folder into the folder itself")
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {

		if f.IsDir() {

			err = CopyDir(src+"/"+f.Name(), dest+"/"+f.Name())
			if err != nil {
				return err
			}

		}

		if !f.IsDir() {

			content, err := ioutil.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err

			}

			err = ioutil.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err

			}

		}

	}
	log.Println("[COPY] Succes to copy file to " + dest)
	return nil
}

func FindIndex(arrs, key interface{}) int {
	var tempArr []string
	tempStr := fmt.Sprint(arrs)
	tempArr = strings.Fields(tempStr[1 : len(tempStr)-1])
	for index, arr := range tempArr {
		if arr == fmt.Sprint(key) {
			return index
		}
	}
	return -1
}

func GetPWDLocation() (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return mydir, nil
}

func CheckIndexNext(arr []string, indexs ...int) bool {
	for _, val := range indexs {
		if val == -1 || val+1 >= len(arr) {
			return false
		}
	}
	return true
}
