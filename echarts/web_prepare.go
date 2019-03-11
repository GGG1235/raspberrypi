package echarts

import (
	"fmt"
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func Service(address string,port int) {
	println(fmt.Sprintf("%s:%d",address,port))
	runtime.GOMAXPROCS(3)
	router := gin.Default()
	router.Use(Cors())
	go func() {
		for {
			times,useds,rams,temps := Prepare(QueryData())
			toUsedHTML(times,useds)
			toTempHTML(times,temps)
			toRamHTML(times,rams)
			toHTML(times,useds,rams,temps)
			time.Sleep(time.Duration(5*60)*time.Second)
		}
	}()
	go router.LoadHTMLGlob("html/*")
	go router.GET("/used",getUsed)
	go router.GET("/temp",getTemp)
	go router.GET("/ram",getRam)
	go router.GET("/",getIndex)
	_ = router.Run(fmt.Sprintf("%s:%d",address,port))
}

func getUsed(c *gin.Context)  {
	c.HTML(http.StatusOK,"used.html",nil)
}

func getTemp(c *gin.Context)  {
	c.HTML(http.StatusOK,"temp.html",nil)
}

func getRam(c *gin.Context)  {
	c.HTML(http.StatusOK,"ram.html",nil)
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK,"lines.html",nil)
}

func toHTML(times,useds,rams,temps interface{}) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title:"CPU_Use"})
	line.AddXAxis(times).
		AddYAxis("CPU_Use",useds)
	line1 := charts.NewLine()
	line1.SetGlobalOptions(charts.TitleOpts{Title:"RAM_Used"})
	line1.AddXAxis(times).
		AddYAxis("RAM_Used",rams)
	line2 := charts.NewLine()
	line2.SetGlobalOptions(charts.TitleOpts{Title:"CPU_Temperature"})
	line2.AddXAxis(times).
		AddYAxis("CPU_Temperature",temps)
	f,err := os.Create("html/lines.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
	line1.Render(f)
	line2.Render(f)
}

func toUsedHTML(times,useds interface{}) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title:"CPU_Use"})
	line.AddXAxis(times).
		AddYAxis("CPU_Use",useds)
		//AddYAxis("RAM_Used",rams)
		//AddYAxis("CPU_Temperature",temps)
	f,err := os.Create("html/used.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}

func toRamHTML(times,rams interface{})  {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title:"RAM_Used"})
	line.AddXAxis(times).
		//AddYAxis("CPU_Use",useds)
	AddYAxis("RAM_Used",rams)
	//AddYAxis("CPU_Temperature",temps)
	f,err := os.Create("html/ram.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}

func toTempHTML(times,temps interface{})  {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title:"CPU_Temperature"})
	line.AddXAxis(times).
		//AddYAxis("CPU_Use",useds)
		//AddYAxis("RAM_Used",rams)
	AddYAxis("CPU_Temperature",temps)
	f,err := os.Create("html/temp.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}
