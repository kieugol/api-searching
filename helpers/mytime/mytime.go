package mytime

import (
	"fmt"
	"time"
)

var loc *time.Location

func SetTimezone(timezone string) error {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	loc = location
	return nil
}

func Now() time.Time {
	fmt.Println(time.Now().In(loc))

	return time.Now().In(loc)
}

func NowUTC() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}
