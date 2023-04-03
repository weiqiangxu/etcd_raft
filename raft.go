package main

type PrometheusHA interface {
	SetPrometheusConfig()       // 设置prometheus的配置
	GetGlobalLockKey()          // 获取全局锁的key
	RegisterNode()              // 注册节点
	RegisterMasterNode()        // 注册主节点
	GetMasterNode()             // 获取主节点
	GetNodeList()               // 获取所有节点列表
	GetConfigYMLString() string // 获取指标采集配置YML格式字符串
	BuildMasterNodeConfig()     // 获取master节点prometheus的配置 - 直接从exporter之中拉取数据
	BuildSlaveNodeConfig()      // 获取slave节点的p的配置 - 从master之中拉取数据
	ConfigReload()              // prometheus配置重载
	StopMasterPullData()        // 主节点停止拉取数据
	CheckSlaveNodeReady()       // 从节点数据拉取完毕状态
}

var _ PrometheusHA = new(PrometheusHATool)

type PrometheusHATool struct {
}

func (p PrometheusHATool) SetPrometheusConfig() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) GetGlobalLockKey() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) RegisterNode() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) RegisterMasterNode() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) GetMasterNode() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) GetNodeList() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) GetConfigYMLString() string {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) BuildMasterNodeConfig() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) BuildSlaveNodeConfig() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) ConfigReload() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) StopMasterPullData() {
	//TODO implement me
	panic("implement me")
}

func (p PrometheusHATool) CheckSlaveNodeReady() {
	//TODO implement me
	panic("implement me")
}

// 主从
// 1. 数据副本通过远程读写 - 需要实现读写适配器
// 2. 联邦机制 - 现有方案，直接使用 - 但是主从变更时候需要重载配置
