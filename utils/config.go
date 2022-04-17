package utils

import (
	"encoding/json"
	"io/ioutil"
	"liereal.com/zinx-kcp/ziface"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数也可以通过 用户根据 zinx.json来配置
*/
type Config struct {
	Server           ziface.IServer //当前Zinx的全局Server对象
	Host             string         //当前服务器主机IP
	Port             int            //当前服务器主机监听端口号
	Name             string         //当前服务器名称
	Version          string         //当前Zinx版本号
	MaxPacketSize    uint32         //当前数据包的最大值
	MaxConn          int            //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32         //当前服务器主机工作池的数量
	MaxWorkerTaskLen uint32
	MaxMsgChanLen    uint32
}

var ConfigInstance *Config

func (c *Config) Reload() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &ConfigInstance)
	if err != nil {
		panic(err)
	}
}

func init() {
	ConfigInstance = &Config{
		Name:             "Zinx-kcp",
		Version:          "0.1",
		Host:             "0.0.0.0",
		Port:             777,
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}

	ConfigInstance.Reload()
}
