package main

import (
        "context"
        "fmt"
        "time"

        "github.com/olivere/elastic"
)

const LOG_INDEX_NAME_FORMAT = "logstash-2006.01.02"

func RemoveLogsBefore(before time.Time, client *elastic.Client) (*elastic.IndicesDeleteResponse, error) {
        names, err := client.IndexNames()
        if err != nil {
                return nil, err
        }

        var needRemoved []string
        for _, name := range names {
                t, err := time.Parse(LOG_INDEX_NAME_FORMAT, name)
                if err != nil {
                        continue
                }
                if t.Before(before) {
                        needRemoved = append(needRemoved, name)
                }
        }

        fmt.Println("try to remove", needRemoved)
        return client.DeleteIndex(needRemoved...).Do(context.Background())
}
