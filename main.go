package main

import (
	"fmt"
	"time"
)

func main() {
	es, err := getEsClient()
	if err != nil {
		fmt.Println("can not connect with elasticsearch", err)
		return
	}

	fmt.Println("cleaner is running")

	for {
		<-time.After(24 * time.Hour)

		t := time.Now().Add(-1 * ONLY_REMAIN_LOG_IN_DURATION)
		_, err := RemoveLogsBefore(t, es)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("cleaned")
	}

}
