package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worklist struct {
	list []string
	Order_ID int
}

func NewWorkList()*Worklist{
	wl := Worklist{Order_ID:0}
	return &wl
}

func (self *Worklist) InputWork(){
	var work string
	for  {
		fmt.Printf("Work (0 to end):")
		fmt.Scanf("%s",&work)
		if work != "0"{
			self.list = append(self.list, work)
			continue
		}
		return
	}
}

func(self *Worklist) GetWork(mode string)string{
	switch mode {
	case "order":
		index := self.Order_ID
		work := self.list[index]
		self.Order_ID = (index+1)%len(self.list)
		return work
	case "random":
		fallthrough
	default:
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(self.list))
		return self.list[index]
	}
}
