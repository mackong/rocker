package main

import (
	"fmt"
	"github.com/RedDragonet/rocker/container"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

func logContainer(containerName string, follow bool) {
	//todo::实现follow

	dirURL := path.Join(container.DefaultInfoLocation, containerName)
	logFileLocation := path.Join(dirURL, container.ContainerLogFile)
	file, err := os.Open(logFileLocation)
	defer file.Close()
	if err != nil {
		log.Errorf("Log container open file %s error %v", logFileLocation, err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("Log container read file %s error %v", logFileLocation, err)
		return
	}
	fmt.Fprint(os.Stdout, string(content))
}
