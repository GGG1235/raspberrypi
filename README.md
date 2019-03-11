# raspberrypi

## 环境:
<img src="https://github.com/GGG1235/raspberrypi/blob/master/images/1.png" width="375" alt="neofetch">

树莓派3b+

go1.12 linux/arm

sqlite3.16.2

---
使用Go语言获得树莓派系统相关信息(cpu占用、温度、RAM使用、磁盘占用),并把获得到的cpu温度、cpu占用率、RAM使用的数据存到sqlite3数据库中,时间间隔为5秒。读取数据库中最近的60组数据,并使用go-echarts生成图像,并且图像在每5分钟更新一次。
在数据存入数据时,遇到一个异常database is locked。
<img src="https://github.com/GGG1235/raspberrypi/blob/master/images/5.png" width="375" alt="error">
网上看到4中解决方案:
>> 1、线程同步
>> 2、写函数对数据库状态进行判断
>> 3、对数据库操作次数延时
>> 4、对数据库操作时间延时


## 图像截图:
<img src="https://github.com/GGG1235/raspberrypi/blob/master/images/2.png" width="375" alt="cpu温度">
cpu温度
<img src="https://github.com/GGG1235/raspberrypi/blob/master/images/3.png" width="375" alt="RAM使用">
RAM使用
<img src="https://github.com/GGG1235/raspberrypi/blob/master/images/4.png" width="375" alt="cpu占用率">
cpu占用率
