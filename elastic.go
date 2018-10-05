package main

import (
	"github.com/olivere/elastic"
)

func getClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(ELASTICSEARCH_URL))
	if err != nil {
		return nil, err
	}
	return client, nil
}