package main

import (
	"flag"
	"fmt"
	"log"
	"raspberrypi/utils"
	"strings"
)

func main(){
	ins := flag.String("ins","","input your instruction")
	flag.Parse()

	if *ins == "" {
		utils.Main()
		return
	}
	res := ""
	Ins := strings.Split(*ins,",")
	i := 0
	for _,p := range Ins {
		switch p {
		case "used":
			res += fmt.Sprintf("CPU = %.2f %%\n",utils.GetCPUuse())
			i++
		case "temp":
			res += fmt.Sprintf("CPU = %.2f ℃\n",utils.GetCPUTemp())
			i++
		case "ram":
			total,used,free,err :=utils.GetRAMInfo()
			if err != nil {
				log.Println(err)
			}
			res += fmt.Sprintf("total = %.2f MB,used = %.2f MB,free = %.2f MB\n",total/1024,used/1024,free/1024)
			i++
		case "disk":
			Size,Used,avail,err := utils.GetDiskSpace()
			if err != nil {
				log.Println(err)
			}
			res +=fmt.Sprintf("size = %s ,used = %s ,avail = %s \n",Size,Used,avail)
			i++
		default:
			continue
		}
	}
	if i == 0 || i>=4 {
		utils.Main()
		return
	} else {
		p := []rune(res)
		fmt.Println(string(p[0 : len(p)-1]))
	}
}
