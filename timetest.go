package main

import (
	"fmt"
	"time"
)

func expensiveCall() {}
func main() {

	//timt := time.Now()
	//
	//fmt.Println(timt.Format("2006-01-02 "))

	//t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Printf("Go launched at %s\n", timt)
	//start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	//end := time.Date(2000, 1, 1, 12, 10, 0, 0, time.UTC)
	//
	//difference := end.Sub(start)
	////fmt.Println(difference)
	//fmt.Printf("difference = %v   %v\n", difference.Seconds() , difference.Nanoseconds())
	//start := time.Now()
	//d, _ := time.ParseDuration("500ms")
	//time.Sleep(d)
	//fmt.Println(time.Since(start).Hours())
	//fmt.Println(time.Until(start))

	//
	//d, err := time.ParseDuration("2h15m30.918273645s")
	//if err != nil {
	//	panic(err)
	//}
	//
	//
	//fmt.Println(d.Round(time.Hour))

	//round := []time.Duration{
	//	time.Nanosecond,
	//	time.Microsecond,
	//	time.Millisecond,
	//	time.Second,
	//	2 * time.Second,
	//	time.Minute,
	//	10 * time.Minute,
	//	time.Hour,
	//}
	//
	//for _, r := range round {
	//	fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	//}

	//d, _ := time.ParseDuration("1h40m40.918273645s")

	//fmt.Println(d.Truncate(time.Minute))
	//if err != nil {
	//	panic(err)
	//}
	//
	//trunc := []time.Duration{
	//	time.Nanosecond,
	//	time.Microsecond,
	//	time.Millisecond,
	//	time.Second,
	//	2 * time.Second,
	//	time.Minute,
	//	10 * time.Minute,
	//	time.Hour,
	//}
	//
	//for _, t := range trunc {
	//	fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	//}

	//	local := time.FixedZone("Beijing", 8*60*60)
	//
	//fmt.Println(local)

	//t, _ := time.LoadLocation("Local")
	//fmt.Println(t.String())

	//now := time.Now()
	//fmt.Println(now)

	//beijng := time.FixedZone("beijing time", 8*60*60)

	//beijing1, _ := time.LoadLocation("Local")
	//a := time.Date(2019, 10, 10, 00, 00, 00, 00, time.UTC)
	//b := time.Date(2019, 10, 10, 00, 00, 00, 00, beijing1)
	//fmt.Println(a)
	//fmt.Println(b)

	//fmt.Printf("%+v\n", a.Date())

	//fmt.Println(a.Month().String())
	//fmt.Println(a.Local())

	//time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "")
	//fmt.Println(time.RFC3339)
	//fmt.Println(time.UnixDate)
	//fmt.Println(time.ANSIC)

	//da, _ := time.Parse("2006-01-02 15:04-07:00", "2019-10-10 08:30+08:00")
	//fmt.Println(da)

	beijng, _ := time.LoadLocation("Local")
	am := time.FixedZone("am", -7*60*60)
	da, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-11-24 08:31:10", beijng)
	da1, _ := time.Parse("2006-01-02 15:04", "2019-10-10 00:30")
	da2, _ := time.ParseInLocation("2006-01-02 15:04", "2019-10-10 00:30", am)

	fmt.Println(da.Zone())
	fmt.Println(da1.String())
	//fmt.Println(da.Date())
	//fmt.Println(da.Day())
	//fmt.Println(da, da1)
	//fmt.Println(da.Hour(), da1.Hour())
	//fmt.Println(da.Equal(da1))
	//fmt.Println(da.Format("2006-01-02 15:04"))
	//fmt.Println(da1.Format("2006-01-02 15:04"))
	//fmt.Printf("%Tv:%v\n", da.Unix(), da.Unix())

	fmt.Println(time.RFC3339)
	fmt.Println(time.UTC)
	fmt.Println(da.Format(time.RFC3339))
	fmt.Println("da round", da.Round(time.Hour))
	fmt.Println("da truncate", da.Truncate(time.Hour))

	fmt.Println(da1.Format(time.RFC3339))
	fmt.Println(da1.IsZero())
	fmt.Println(da.IsZero())
	fmt.Println("da1 local", da1.Local())
	fmt.Println(da2.Format(time.RFC3339))

	//da1.UnmarshalJSON()

	var da3 time.Time
	da3.UnmarshalText([]byte("2019-11-24T08:30:00+08:00"))
	fmt.Println("Da3", da3)

	var da4 time.Time
	da4.UnmarshalBinary([]byte("2019-11-24T08:30:00+08:00"))
	fmt.Println("da4", da4)

	da5, _ := time.Parse("2006-01-02 00:00-07:00", "2019-01-30 00:00-07:00")
	fmt.Println(da5.ISOWeek())

	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))
	//fmt.Println(time.Unix(20, 1570667400*1000000000))

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	da6, _ := time.Parse("2006-01-02 15:04", "2019-11-24 23:30")
	fmt.Println(da6)

	localTz, _ := time.LoadLocation("Local")
	da7, _ := time.ParseInLocation("2006-01-02 15:04", "2019-11-24 23:30", localTz)
	fmt.Println(da7)

	fmt.Println("-=------")

	fmt.Println(time.Now().Add(time.Second * 60))

	fmt.Println(time.Now().Add(time.Hour))
	fmt.Println(time.Now().Add(time.Minute))

	fmt.Println(time.Now().AddDate(1, 0, 0))
	fmt.Println(time.Now().AddDate(0, 1, 0))

	//time.Unix()

	fmt.Println(time.Now().Unix())

	fmt.Println(time.Unix(1574608496, 0))



	fmt.Println(time.Now().Truncate(time.Minute))
	fmt.Println(time.Now().Truncate(time.Hour))

	//day, _ := time.ParseDuration("24h")
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Truncate(day).Add(time.Hour*-8))




	//24小时
	//da8, _ := time.ParseInLocation("2006-01-02 15:04", "2019-11-20 10:10", localTz)
	//
	//fmt.Println(time.Now().Sub(da8))

}
