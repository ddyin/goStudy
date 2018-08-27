package main

import (
	"fmt"
	"time"
)

//时间
func main()  {
	p:=fmt.Println
	now :=time.Now()

	p(now)

	then :=time.Date(2018, 8, 16, 9, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff :=now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

}

