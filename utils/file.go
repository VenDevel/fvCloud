package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFileAllByPath(path string) ([]byte, error) {
	fi, err := os.Open(path)
	log.Println(path)
	if err != nil {
		log.Println("----------------------")
		log.Println(path, err.Error())
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd, err
}
