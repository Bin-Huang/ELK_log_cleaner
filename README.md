## log cleaner service

可以帮你定时清理 ELK 中的过时日志，防止日志占满磁盘。运行内存约 `1.6mb`。

默认每天清理已保存超过一个月的日志，可以修改 [config.go](./config.go) 进行配置
