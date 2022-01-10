package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func IsExist(fp string) bool  {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

func ToTrimString(filePath string) (string, error){
	str, err := toString(filePath)
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(str),nil
}

func toString(filePath string) (string, error)  {
	b, err := ioutil.ReadFile(filePath)
	if err != nil{
		return "", err
	}
	return string(b), nil

}