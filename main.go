package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"strconv"
	"time"
)

func PickUnUsedPort()(int, error){
	addr, err :=  net.ResolveTCPAddr("tcp","127.0.0.1:")
	if err != nil {
		return 0, err
	}
	l ,err :=  net.ListenTCP("tcp",addr)
	if err != nil {
		return 0,err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0,err
	}
	return port, nil
}

func main(){
	port,err := PickUnUsedPort()
	if err != nil {
		panic("Get Port Failed !!")
	}
	fmt.Println("port:",port)
	opts := []selenium.ServiceOption{
		// Enable fake XWindow session.
		// selenium.StartFrameBuffer(),
		//selenium.Output(os.Stderr), // Output debug information to STDERR
	}
	var user Info
	fmt.Printf("UserName:")
	fmt.Scanf("%s",&user.UserName)
	fmt.Printf("PassWord:")
	fmt.Scanf("%s",&user.PassWord)

	var Workinfo WorkTime
	fmt.Printf("Year:")
	fmt.Scanf("%d",&Workinfo.Year)
	fmt.Printf("Month:")
	fmt.Scanf("%d",&Workinfo.Month)
	fmt.Printf("WorkHour:")
	fmt.Scanf("%d",&Workinfo.WorkHour)
	hourZone, Lefthour := Workinfo.WorkHour/4, Workinfo.WorkHour%4

	workday := AppendWorkday(Workinfo.Year,Workinfo.Month)
	worklist := NewWorkList()
	worklist.InputWork()

	selenium.SetDebug(false)

	service := NewService(port,opts)
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	driver := RemoteService(port,caps)
	defer driver.Quit()

	GetPage(driver,"https://miswww1.ccu.edu.tw/parttime/index.php")
	LoginSystem(driver,user)
	if ok,err := IsLogin(driver);!ok {
		panic(err)
	}
	fmt.Println("Login Success!!")

	Year := strconv.Itoa(Workinfo.Year)
	Month := strconv.Itoa(Workinfo.Month)

	for i:=0;i<=hourZone;i++{
		if i == hourZone && Lefthour == 0 {
			break
		}
		switch state:=i%2;state{
		case 0:
			if i != hourZone {
				InputOneRoutine(driver, Year, Month, <-workday, worklist.GetWork("order"), "08", "12")
			} else {
				InputOneRoutine(driver, Year, Month, <-workday, worklist.GetWork("order"), "08", fmt.Sprintf("%d",8+Lefthour))
			}
		case 1:
			if i != hourZone {
			InputOneRoutine(driver,Year,Month,<-workday,worklist.GetWork("order"),"13","17")
			} else {
				InputOneRoutine(driver, Year, Month, <-workday, worklist.GetWork("order"), "13", fmt.Sprintf("%d",13+Lefthour))
			}
		}
	}

	GetPage(driver,"https://miswww1.ccu.edu.tw/parttime/control2.php")
	GetPage(driver,"https://miswww1.ccu.edu.tw/parttime/main2.php")
	ClickAction(driver,"/html/body/form/center/input[3]")
	time.Sleep(30* time.Second)
}