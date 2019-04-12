package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func GetPage(driver selenium.WebDriver,url string){
	for {
		err := driver.Get(url)
		if err != nil {
			fmt.Println("Get Page Failed !!, retrying...")
			continue
		}
		return
	}
}

func Inputinfo(driver selenium.WebDriver,elemValue, input_info string)error {
	elem,err := driver.FindElement(selenium.ByName, elemValue)
	err = elem.SendKeys(fmt.Sprintf("%s%s%s",selenium.BackspaceKey,selenium.BackspaceKey,selenium.BackspaceKey))
	err = elem.SendKeys(input_info)
	return err
}

func ClickAction(driver selenium.WebDriver,elemValue string)error {
	elem,err := driver.FindElement(selenium.ByXPATH,elemValue)
	err = elem.Click()
	return err
}

func InputOneRoutine(driver selenium.WebDriver,year,month,day,work,shour,ehour string) {
	GetPage(driver,"https://miswww1.ccu.edu.tw/parttime/main2.php")
	Inputinfo(driver,"yy",year)
	Inputinfo(driver,"mm",month)
	Inputinfo(driver,"dd",day)
	Inputinfo(driver,"workin",work)
	Inputinfo(driver,"shour",shour)
	Inputinfo(driver,"ehour",ehour)
	ClickAction(driver,`//select[@name="type"]/option[@value="Z135奈米生物檢測科技研究中心"]`)
	ClickAction(driver,"/html/body/form/center/input[2]")
}