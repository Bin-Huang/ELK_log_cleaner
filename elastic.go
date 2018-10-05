package main

import (
	"github.com/olivere/elastic"
)

func getClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}
	return client, nil
}