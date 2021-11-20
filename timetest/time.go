package timetest

import (
	"fmt"
	"time"
)

func Timetest(){
	t:=time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())
	//
	t=time.Now().UTC()
	fmt.Println(t)
	fmt.Print(time.Now())
	//计算时间
	var week time.Duration
	week=60*60*24*7*1e9 //must be in nanosec
	weekFromNow:=t.Add(week)
	fmt.Println(weekFromNow)
	//格式化时间
	fmt.Println(t.Format(time.RFC822)) //21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s:=t.Format("20060102")
	fmt.Println(t,"=>",s)

}
