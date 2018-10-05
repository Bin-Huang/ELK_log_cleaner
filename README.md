## log cleaner service

定时自动清理 ELK 的过时日志，防止日志占满磁盘。

默认每天清理已保存超过一个月的日志，可以修改 [config.go](./config.go) 进行配置