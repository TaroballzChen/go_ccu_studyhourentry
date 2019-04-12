package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func NewService(port int, opts []selenium.ServiceOption)(service *selenium.Service){
	service, err := selenium.NewChromeDriverService("./chromedriver", port, opts...)
	if err != nil {
		panic(err)
	}

	return
}

func RemoteService(port int,caps selenium.Capabilities)(wd selenium.WebDriver){
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	return
}