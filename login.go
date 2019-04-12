package main

import (
	"errors"
	"github.com/tebeka/selenium"
	"strings"
)

type Info struct {
	UserName string
	PassWord string
}


func LoginSystem(driver selenium.WebDriver,UserInfo Info)error{
	UserNameElem, err := driver.FindElement(selenium.ByName,"staff_cd")
	UserNameElem.SendKeys(UserInfo.UserName)
	PassWdElem, err := driver.FindElement(selenium.ByName,"passwd")
	PassWdElem.SendKeys(UserInfo.PassWord)
	ButtonElem, err := driver.FindElement(selenium.ByXPATH,"/html/body/center/form/input[1]")
	ButtonElem.Click()

	return err
}

func IsLogin(driver selenium.WebDriver)(bool,error) {
	driver.Get("https://miswww1.ccu.edu.tw/parttime/frame_stu.php")
	centerElem,err := driver.FindElement(selenium.ByXPATH, "/html/body/center")
	if ok := strings.Contains(err.Error(),"no such element");ok {
		return true,nil
	}
	text,err := centerElem.Text()
	if text[:15] == "請正常登入" {
		return false,errors.New("Login Failed!!")
	}
	return false,err
}