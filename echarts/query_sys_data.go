package echarts

import (
	"log"
	"raspberrypi/utils"
)

func QueryData() (res []map[string]interface{}){
	database := utils.NewSqlite("./sys_info")
	db,err := database.OpenSource()
	if err != nil {
		log.Println(err)
	}
	res = database.QueryData(db,"SELECT * FROM sys_info order by strftime desc LIMIT 45")
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return
}

func Prepare(data []map[string]interface{}) (x,useds,rams,temps []interface{}) {
	x = make([]interface{},0)
	useds = make([]interface{},0)
	rams = make([]interface{},0)
	temps = make([]interface{},0)
	for _,item := range data {
		x=append(x,item["strftime"])
		useds=append(useds,item["used"])
		rams=append(rams,item["ram"])
		temps=append(temps,item["temp"])
	}
	return
}