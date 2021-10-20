package helper

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFileArray(path, seperator string) ([]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	temp := strings.Split(string(dat), seperator)
	log.Println("[READ] Success to read file " + path)
	return temp[:len(temp)-1], nil
}
func GetValueByName(path, name string) ([]string, error) {
	fmt.Println(name)
	arrs, err := ReadFileArray(path, "\n")
	if err != nil {
		return nil, nil
	}
	for _, arr := range arrs {
		split := strings.Split(arr, "#")
		if split[0] == name {
			return split, nil
		}
	}
	return nil, nil
}
func ReadFileString(path, seperator string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	log.Println("[READ] Success to read file " + path)
	return string(dat), nil
}

func ReadFileScan(path, seperator string) ([]string, error) {
	var temp string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		temp += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	log.Println("[READ] Success to read file " + path)
	return strings.Split(temp, seperator), nil
}

func AppendToFile(path, text string) error {
	currentValue, err := ReadFileArray(path, "\n")

	if err == nil {
		for _, field := range currentValue {
			fields := strings.Split(field, "#")
			currentfields := strings.Split(text, "#")
			if fields[0] == currentfields[0] {
				log.Println("[DUPLICATE] Sorry name already exist")
				return errors.New("sorry name already exist")
			}
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	log.Println("[WRITE] Success to append string")
	return nil
}
func WriteToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	log.Println("[WRITE] Success to append string")
	return nil
}

func DeleteFileWithSeperator(path, nameDelete, seperator string) error {
	tempFile, err := ReadFileString(path, seperator)
	if err != nil {
		return err
	}
	tempsFile := strings.Split(tempFile, "\n")
	tempsFile = tempsFile[:len(tempsFile)-1]
	newStr := ""
	for _, field := range tempsFile {
		fields := strings.Split(field, "#")
		if fields[0] != nameDelete {
			newStr += field + "\n"
		}
	}

	if tempFile != newStr {
		err = WriteToFile(path, newStr)
		if err != nil {
			return err
		}
		log.Println("[Delete] Success delete " + nameDelete + " from " + path)
	} else {
		log.Println("[Delete] Name " + nameDelete + " not found")
	}
	return nil
}

func EditFileWithSeperator(path, nameEdit, newValue, seperator string) error {
	tempFile, err := ReadFileString(path, seperator)
	found := false
	// fmt.Println(tempFile)
	if err != nil {
		return err
	}
	tempsFile := strings.Split(tempFile, "\n")
	tempsFile = tempsFile[:len(tempsFile)-1]
	newStr := ""
	for _, field := range tempsFile {
		fields := strings.Split(field, "#")
		// fmt.Print(field)
		if fields[0] != nameEdit {
			newStr += field + "\n"
		} else {
			found = true
			newStr += fields[0] + "#" + newValue + "\n"
		}
	}
	if tempFile != newStr {
		err = WriteToFile(path, newStr)

		if err != nil {
			return err
		}
		log.Println("[Edit] Success edit " + nameEdit + " from " + path)
	} else if found {
		log.Println("[Edit] Name " + nameEdit + " found but not change because value same")
	} else {
		log.Println("[Edit] Name " + nameEdit + " not found")
	}
	return nil
}
