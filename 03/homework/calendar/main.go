package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Calendar struct {
	WeekLabels   []string
	MonthLabels  []string
	Date         time.Time
	IsDayOffFunc func(t time.Time) bool
	LinkFunc     func(t time.Time) string
	SelectedDate time.Time
}

var DefaultWeekLabels = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func DefaultIsDayOffFunc(t time.Time) bool {
	return t.Weekday() == time.Sunday || t.Weekday() == time.Saturday
}

func NewCalendar() *Calendar {
	return &Calendar{WeekLabels: DefaultWeekLabels, IsDayOffFunc: DefaultIsDayOffFunc, Date: time.Now()}
}

func (c *Calendar) PrintCalendar() string {
	wc := len(c.WeekLabels)
	wd := (int(c.Date.Weekday()) - (c.Date.Day() - 1) + wc*30) % wc
	last := c.Date.AddDate(0, 1, -c.Date.Day()).Day()

	current := time.Now().UTC()
	y, m, _ := current.Date()
	header := m.String() + " " + strconv.Itoa(y)
	fmt.Println(fmt.Sprintf("%[1]*s", -35, fmt.Sprintf("%[1]*s", (35 + len(header))/2, header)))
	s := " " + strings.Join(c.WeekLabels, "  ") + " \n"
	//s += "|" + strings.Repeat("----:|", wc) + "\n"
	s += "" + strings.Repeat("     ", wd)
	for d := 1; d <= last; d++ {
		//url := c.link(d)
		//if url != "" {
		//	s += " [" + fmt.Sprint(d) + "](" + url + ") |"
		//} else {
		s += fmt.Sprintf("  %2d ", d)
		//}
		wd = (wd + 1) % wc
		if wd == 0 {
			s += "\n"
			if d != last {
				s += ""
			}
		}
	}
	s += strings.Repeat("     ", wc-wd)

	return s
}

func (c *Calendar) String() string {
	return c.PrintCalendar()
}

//func (c Calendar) link(d int) string {
//	if c.LinkFunc == nil {
//		return ""
//	}
//	return c.LinkFunc(c.Date.AddDate(0, 0, -c.Date.Day()+d))
//}

func main()  {
	fmt.Println(NewCalendar().PrintCalendar())
}