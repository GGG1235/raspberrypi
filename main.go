package main

import (
	"fmt"
	"raspberrypi/utils"
	"time"
)

func main(){
	database := utils.NewSqlite("./sys_info")
	db,_ := database.OpenSource()
	err:=database.CreateDataTable(db)
	if err != nil {
		fmt.Println(err)
	}
	func() {
		for {
			temp:=utils.GetCPUTemp()
			used:=utils.GetCPUuse()
			_,ram,_,_:=utils.GetRAMInfo()
			go database.InsertData(db,temp,used,ram/1024)
			time.Sleep(time.Duration(5)*time.Second)
		}
	}()
	database.CloseConn(db)
}
