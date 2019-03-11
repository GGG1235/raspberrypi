package main

import (
	"flag"
	"raspberrypi/echarts"
)

func main() {
	address := flag.String("addr","127.0.0.1","input address")
	port := flag.Int("port",8080,"input port")
	if *port >= 65535 {
		*port = 8080
	}
	flag.Parse()
	echarts.Service(*address,*port)
}
