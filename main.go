package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AndreeJait/bagong.git/helper"
)

func main() {
	cmdArgs := os.Args[1:]
	baseBagong, err := helper.GetBaseBagong()
	if err != nil {
		var tempBase string
		fmt.Print("Input base bagong : ")
		fmt.Scan(&tempBase)
		err := os.Setenv("BASE_BAGONG", tempBase)
		if err != nil {
			log.Fatalln(err)
		}
		baseBagong = os.Getenv("BASE_BAGONG")
		fmt.Println("Success to set base bagong")
	}
	currentLocation, err := helper.GetPWDLocation()

	if err != nil {
		log.Fatalln(err)
	}

	if cmdArgs[0] == "template" {
		baseFile := baseBagong + "/config/template_config.txt"
		if cmdArgs[1] == "add" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name")
			indexValueTemplate := helper.FindIndex(cmdArgs, "--value")
			if helper.CheckIndexNext(cmdArgs, indexNameTemplate, indexValueTemplate) {
				err := helper.AppendToFile(baseFile, cmdArgs[indexNameTemplate+1]+"#"+currentLocation+cmdArgs[indexValueTemplate+1]+"\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				} else {
					log.Println("[Success] success to add template")
				}

			}
		} else if cmdArgs[1] == "list" {
			list, err := helper.ReadFileArray(baseFile, "\n")
			if err != nil {
				log.Println(err)
			}
			fmt.Println("NAME#VALUE")
			for _, l := range list {
				fmt.Println(l)
			}
		} else if cmdArgs[1] == "get" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			indexFolderName := helper.FindIndex(cmdArgs, "--folder") + 1
			indexDestination := helper.FindIndex(cmdArgs, "--dest") + 1

			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1, indexFolderName-1, indexDestination-1) {
				templateInfo, err := helper.GetValueByName(baseFile, cmdArgs[indexNameTemplate])
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
				err = helper.CopyDir(templateInfo[1]+"/"+cmdArgs[indexFolderName], currentLocation+"/"+cmdArgs[indexDestination])

				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		} else if cmdArgs[1] == "edit" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			indexValueName := helper.FindIndex(cmdArgs, "--value") + 1

			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1, indexValueName-1) {
				err := helper.EditFileWithSeperator(baseFile, cmdArgs[indexNameTemplate], currentLocation+"/"+cmdArgs[indexValueName], "\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		} else if cmdArgs[1] == "delete" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1) {
				err := helper.DeleteFileWithSeperator(baseFile, cmdArgs[indexNameTemplate], "\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		}
	} else if cmdArgs[0] == "resource" {
		baseFile := baseBagong + "/config/resource_config.txt"
		if cmdArgs[1] == "add" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name")
			indexValueTemplate := helper.FindIndex(cmdArgs, "--value")
			if helper.CheckIndexNext(cmdArgs, indexNameTemplate, indexValueTemplate) {
				err := helper.AppendToFile(baseFile, cmdArgs[indexNameTemplate+1]+"#"+cmdArgs[indexValueTemplate+1]+"\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				} else {
					log.Println("[Success] success to add template")
				}

			}
		} else if cmdArgs[1] == "list" {
			list, err := helper.ReadFileArray(baseFile, "\n")
			if err != nil {
				log.Println(err)
			}
			fmt.Println("NAME\tVALUE")
			for _, l := range list {
				split := strings.Split(l, "#")
				fmt.Println(split[0] + "\t" + split[1])
			}
		} else if cmdArgs[1] == "get" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			indexDestination := helper.FindIndex(cmdArgs, "--dest") + 1

			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1, indexDestination-1) {
				templateInfo, err := helper.GetValueByName(baseFile, cmdArgs[indexNameTemplate])
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
				_, err = helper.CloneGit(templateInfo[1], currentLocation+"/"+cmdArgs[indexDestination])
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
				err = os.RemoveAll(currentLocation + "/" + cmdArgs[indexDestination] + "/.git")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		} else if cmdArgs[1] == "edit" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			indexValueName := helper.FindIndex(cmdArgs, "--value") + 1

			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1, indexValueName-1) {
				err := helper.EditFileWithSeperator(baseFile, cmdArgs[indexNameTemplate], cmdArgs[indexValueName], "\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		} else if cmdArgs[1] == "delete" {
			indexNameTemplate := helper.FindIndex(cmdArgs, "--name") + 1
			if helper.CheckIndexNext(cmdArgs, indexNameTemplate-1) {
				err := helper.DeleteFileWithSeperator(baseFile, cmdArgs[indexNameTemplate], "\n")
				if err != nil {
					log.Println("[ERROR] " + err.Error())
				}
			}
		}
	}
}
