package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Main(){
	fmt.Println(fmt.Sprintf("CPU = %.2f %%",GetCPUuse()))
	fmt.Println(fmt.Sprintf("CPU = %.2f ℃",GetCPUTemp()))
	total,used,free,err :=GetRAMInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(fmt.Sprintf("total = %.2f MB,used = %.2f MB,free = %.2f MB",total/1024,used/1024,free/1024))
	Size,Used,avail,err := GetDiskSpace()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(fmt.Sprintf("size = %s ,used = %s ,avail = %s ",Size,Used,avail))
}

func GetCPUTemp() float64  {
	b,err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp") ///sys/class/thermal/thermal_zone0/temp
	//b,err := ioutil.ReadFile("./utils/temp") ///sys/class/thermal/thermal_zone0/temp
	if err != nil{
		log.Println(err)
	}
	str := string(b)
	rs := []rune(str)
	length := len(rs)
	res,_ :=strconv.ParseFloat(string(rs[0:length-1]),64)

	return float64(res / 1000)
}

func GetCPUuse() float64 {
	c := exec.Command("top","-bn","1","-i","-c") //top -bn 1 -i -c
	//c := exec.Command("tail","./utils/test") //top -bn 1 -i -c
	buf,_ := c.Output()
	str := string(buf)
	reg := regexp.MustCompile(`(\d+\.\d+)`)
	t :=""
	for _,param :=range reg.FindStringSubmatch(strings.Split(str,"\n")[2]) {
		t = param
	}
	res,_ := strconv.ParseFloat(t,64)
	return float64(res)
}

func GetRAMInfo() (float64,float64,float64,error){
	c := exec.Command("free")
	//c := exec.Command("tail","./utils/ram")
	buf,_ := c.Output()
	str := string(buf)
	reg := regexp.MustCompile(`\d+`)
	i := 0
	var lis []float32
	for _,param :=range reg.FindAllString(strings.Split(str,"\n")[1],-1) {
		t,_ := strconv.ParseFloat(param,64)
		lis=append(lis, float32(t))
		i += 1
		if i == 3 {
			break
		}
	}
	if len(lis) != 3 {
		return 0, 0, 0, errors.New("GetRAMInfo failed")
	}
	return float64(lis[0]), float64(lis[1]), float64(lis[2]),nil
}

func GetDiskSpace() (string,string,string,error) {
	c := exec.Command("df","-h","/")
	//c := exec.Command("tail","./utils/disk")
	buf,_ := c.Output()
	str := string(buf)

	reg := regexp.MustCompile(`(\d+G)|(\d\.\d+G)`)
	i := 0
	var lis []string
	for _,param :=range reg.FindAllString(strings.Split(str,"\n")[1],-1) {
		lis=append(lis, param)
		i += 1
		if i == 3 {
			break
		}
	}
	if len(lis) != 3 {
		return "", "", "", errors.New("GetDiskSpace failed")
	}
	return lis[0],lis[1],lis[2],nil
}
