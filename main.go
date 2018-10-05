package main

import (
	"sync"
	"time"
	"fmt"
	"github.com/robfig/cron"
)

var w sync.WaitGroup

func main() {
	cilent, err := getClient()
	if err != nil {
		fmt.Println("can not connect with elasticsearch", err)
		return
	}
	w.Add(1)

	c := cron.New()
	c.AddFunc("@daily", func() {
		t := time.Now().Add(-1 * 30 * 24 * time.Hour)
		_, err := RemoveLogsBefore(t, cilent)
		if err != nil {
			fmt.Println(err)
		}
	})

	c.Start()
	w.Wait()
}