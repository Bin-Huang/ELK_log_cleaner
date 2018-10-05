package main

import "time"

const LOG_INDEX_NAME_FORMAT = "logstash-2006.01.02"	// es 中日志 indices 名称格式

const ELASTICSEARCH_URL = "http://127.0.0.1:9200"	// elasticsearch service url

const ONLY_REMAIN_LOG_IN_DURATION = 30 * 24 * time.Hour	// 仅仅保留近一个月的日志
