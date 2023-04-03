package main

type PrometheusHA interface {
	KeepAlive()             // 定时任务隔10s扫一下prometheus主节点是否存活
	BuildMasterNodeConfig() // 获取master节点prometheus的配置 - 直接从exporter之中拉取数据
	BuildSlaveNodeConfig()  // 获取slave节点的p的配置 - 从master之中拉取数据
}

// 主从
// 1. 数据副本通过远程读写 - 需要实现读写适配器
// 2. 联邦机制 - 现有方案，直接使用 - 但是主从变更时候需要重载配置
